package main

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	hasherFactory "github.com/ElrondNetwork/elrond-go-core/hashing/factory"
	marshalFactory "github.com/ElrondNetwork/elrond-go-core/marshal/factory"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	nodeFactory "github.com/ElrondNetwork/elrond-go/cmd/node/factory"
	"github.com/ElrondNetwork/elrond-go/common/factory"
	"github.com/ElrondNetwork/elrond-go/common/logging"
	erdConfig "github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/ElrondNetwork/elrond-proxy-go/api"
	"github.com/ElrondNetwork/elrond-proxy-go/config"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
	"github.com/ElrondNetwork/elrond-proxy-go/metrics"
	"github.com/ElrondNetwork/elrond-proxy-go/observer"
	"github.com/ElrondNetwork/elrond-proxy-go/process"
	"github.com/ElrondNetwork/elrond-proxy-go/process/cache"
	"github.com/ElrondNetwork/elrond-proxy-go/process/database"
	processFactory "github.com/ElrondNetwork/elrond-proxy-go/process/factory"
	"github.com/ElrondNetwork/elrond-proxy-go/rosetta"
	"github.com/ElrondNetwork/elrond-proxy-go/testing"
	versionsFactory "github.com/ElrondNetwork/elrond-proxy-go/versions/factory"
	"github.com/urfave/cli"
)

const (
	defaultLogsPath      = "logs"
	logFilePrefix        = "elrond-proxy"
	logFileLifeSpanInSec = 86400
	logFileMaxSizeInMB   = 1024
)

var (
	memoryBallastObject []byte
	proxyHelpTemplate   = `NAME:
   {{.Name}} - {{.Usage}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}
VERSION:
   {{.Version}}
   {{end}}
`

	log = logger.GetOrCreate("proxy")

	// profileMode defines a flag for profiling the binary
	// If enabled, it will open the pprof routes over the default gin rest webserver.
	// There are several routes that will be available for profiling (profiling can be analyzed with: go tool pprof):
	//  /debug/pprof/ (can be accessed in the browser, will list the available options)
	//  /debug/pprof/goroutine
	//  /debug/pprof/heap
	//  /debug/pprof/threadcreate
	//  /debug/pprof/block
	//  /debug/pprof/mutex
	//  /debug/pprof/profile (CPU profile)
	//  /debug/pprof/trace?seconds=5 (CPU trace) -> being a trace, can be analyzed with: go tool trace
	// Usage: go tool pprof http(s)://ip.of.the.server/debug/pprof/xxxxx
	profileMode = cli.BoolFlag{
		Name: "profile-mode",
		Usage: "Boolean option for enabling the profiling mode. If set, the /debug/pprof routes will be available " +
			"on the node for profiling the application.",
	}
	// configurationFile defines a flag for the path to the main toml configuration file
	configurationFile = cli.StringFlag{
		Name:  "config",
		Usage: "The main configuration file to load",
		Value: "./config/config.toml",
	}
	// walletKeyPemFile represents the path of the wallet (address) pem file
	walletKeyPemFile = cli.StringFlag{
		Name:  "pem-file",
		Usage: "This represents the path of the walletKey.pem file",
		Value: "./config/walletKey.pem",
	}
	// externalConfigFile defines a flag for the path to the external toml configuration file
	externalConfigFile = cli.StringFlag{
		Name: "config-external",
		Usage: "The path for the external configuration file. This TOML file contains" +
			" external configurations such as ElasticSearch's URL and login information",
		Value: "./config/external.toml",
	}

	// credentialsConfigFile defines a flag for the path to the credentials toml configuration file
	credentialsConfigFile = cli.StringFlag{
		Name: "config-credentials",
		Usage: "The path for the credentials configuration file. This TOML file contains" +
			" a list of username-password pairs able to perform actions on some endpoints.",
		Value: "./config/apiConfig/credentials.toml",
	}

	// apiConfigDirectory defines a flag for the path to the api configuration directory
	apiConfigDirectory = cli.StringFlag{
		Name: "api-config-directory",
		Usage: "The path for the credentials configuration file. This TOML file contains" +
			" a list of username-password pairs able to perform actions on some endpoints.",
		Value: "./config/apiConfig",
	}

	// testHttpServerEn used to enable a test (mock) http server that will handle all requests
	testHttpServerEn = cli.BoolFlag{
		Name:  "test-mode",
		Usage: "Enables a test http server that will handle all requests",
	}

	startAsRosetta = cli.BoolFlag{
		Name:  "rosetta",
		Usage: "Starts the proxy as a rosetta server",
	}
	rosettaOffline = cli.BoolFlag{
		Name:  "offline",
		Usage: "Starts rosetta server offline",
	}
	rosettaOfflineConfig = cli.StringFlag{
		Name:  "offline-config",
		Usage: "The path for the offline rosetta configuration",
		Value: "./../../rosetta/offline_config_devnet.toml",
	}

	// logLevel defines the logger level
	logLevel = cli.StringFlag{
		Name: "log-level",
		Usage: "This flag specifies the logger `level(s)`. It can contain multiple comma-separated value. For example" +
			", if set to *:INFO the logs for all packages will have the INFO level. However, if set to *:INFO,api:DEBUG" +
			" the logs for all packages will have the INFO level, excepting the api package which will receive a DEBUG" +
			" log level.",
		Value: "*:" + logger.LogInfo.String(),
	}
	// logFile is used when the log output needs to be logged in a file
	logSaveFile = cli.BoolFlag{
		Name:  "log-save",
		Usage: "Boolean option for enabling log saving. If set, it will automatically save all the logs into a file.",
	}
	// workingDirectory defines a flag for the path for the working directory.
	workingDirectory = cli.StringFlag{
		Name:  "working-directory",
		Usage: "This flag specifies the `directory` where the proxy will store logs.",
		Value: "",
	}
	// memBallast defines a flag that specifies the number of MegaBytes to be used as a memory ballast for Garbage Collector optimization
	// if set to 0, the memory ballast won't be used
	memBallast = cli.Uint64Flag{
		Name:  "mem-ballast",
		Value: 0,
		Usage: "Flag that specifies the number of MegaBytes to be used as a memory ballast for Garbage Collector optimization. " +
			"If set to 0, the feature will be disabled",
	}

	testServer *testing.TestHttpServer
)

func main() {
	removeLogColors()

	app := cli.NewApp()
	cli.AppHelpTemplate = proxyHelpTemplate
	app.Name = "Elrond Node Proxy CLI App"
	app.Version = "v1.0.0"
	app.Usage = "This is the entry point for starting a new Elrond node proxy"
	app.Flags = []cli.Flag{
		configurationFile,
		externalConfigFile,
		credentialsConfigFile,
		apiConfigDirectory,
		profileMode,
		walletKeyPemFile,
		testHttpServerEn,
		startAsRosetta,
		rosettaOffline,
		rosettaOfflineConfig,
		logLevel,
		logSaveFile,
		workingDirectory,
		memBallast,
	}
	app.Authors = []cli.Author{
		{
			Name:  "The Elrond Team",
			Email: "contact@elrond.com",
		},
	}

	app.Action = startProxy

	defer func() {
		if testServer != nil {
			testServer.Close()
		}
	}()

	err := app.Run(os.Args)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func initializeLogger(ctx *cli.Context) (nodeFactory.FileLoggingHandler, error) {
	logLevelFlagValue := ctx.GlobalString(logLevel.Name)
	err := logger.SetLogLevel(logLevelFlagValue)
	if err != nil {
		return nil, err
	}
	workingDir := getWorkingDir(ctx, log)

	var fileLogging nodeFactory.FileLoggingHandler
	withLogFile := ctx.GlobalBool(logSaveFile.Name)
	if withLogFile {
		fileLogging, err = logging.NewFileLogging(logging.ArgsFileLogging{
			WorkingDir:      workingDir,
			DefaultLogsPath: defaultLogsPath,
			LogFilePrefix:   logFilePrefix,
		})
		if err != nil {
			return nil, fmt.Errorf("%w creating a log file", err)
		}
	}

	if !check.IfNil(fileLogging) {
		err = fileLogging.ChangeFileLifeSpan(time.Second*time.Duration(logFileLifeSpanInSec), logFileMaxSizeInMB)
		if err != nil {
			return nil, err
		}
	}

	return fileLogging, nil
}

func startProxy(ctx *cli.Context) error {
	memBallastValue := ctx.GlobalUint64(memBallast.Name)
	if memBallastValue > 0 {
		// memory ballast is an optimization for golang's garbage collector. If set to a high value, it can decrease
		// the number of times when GC performs STW processes, that results is a better performance over high load
		memoryBallastObject = make([]byte, memBallastValue*core.MegabyteSize)
		log.Info("initialized memory ballast object", "size", core.ConvertBytes(uint64(len(memoryBallastObject))))
	}

	fileLogging, err := initializeLogger(ctx)
	if err != nil {
		return err
	}

	isProfileModeActivated := ctx.GlobalBool(profileMode.Name)

	log.Info("Starting proxy...")

	configurationFileName := ctx.GlobalString(configurationFile.Name)
	generalConfig, err := loadMainConfig(configurationFileName)
	if err != nil {
		return err
	}
	log.Info(fmt.Sprintf("Initialized with main config from: %s", configurationFile))

	externalConfigurationFileName := ctx.GlobalString(externalConfigFile.Name)
	externalConfig, err := loadExternalConfig(externalConfigurationFileName)
	if err != nil {
		return err
	}

	closableComponents := data.NewClosableComponentsHandler()

	credentialsConfigurationFileName := ctx.GlobalString(credentialsConfigFile.Name)
	credentialsConfig, err := loadCredentialsConfig(credentialsConfigurationFileName)
	if err != nil {
		return err
	}

	statusMetricsProvider := metrics.NewStatusMetrics()

	versionsRegistry, err := createVersionsRegistryTestOrProduction(ctx, generalConfig, configurationFileName, externalConfig, statusMetricsProvider, closableComponents)
	if err != nil {
		return err
	}

	httpServer, err := startWebServer(versionsRegistry, ctx, generalConfig, *credentialsConfig, statusMetricsProvider, isProfileModeActivated)
	if err != nil {
		return err
	}

	waitForServerShutdown(httpServer, closableComponents)

	log.Debug("closing proxy")
	if !check.IfNil(fileLogging) {
		err = fileLogging.Close()
		log.LogIfError(err)
	}

	return nil
}

func loadMainConfig(filepath string) (*config.Config, error) {
	cfg := &config.Config{}
	err := core.LoadTomlFile(cfg, filepath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func loadExternalConfig(filepath string) (*erdConfig.ExternalConfig, error) {
	cfg := &erdConfig.ExternalConfig{}
	err := core.LoadTomlFile(cfg, filepath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func createVersionsRegistryTestOrProduction(
	ctx *cli.Context,
	cfg *config.Config,
	configurationFilePath string,
	exCfg *erdConfig.ExternalConfig,
	statusMetricsHandler data.StatusMetricsProvider,
	closableComponents *data.ClosableComponentsHandler,
) (data.VersionsRegistryHandler, error) {

	var testHTTPServerEnabled bool
	if ctx.IsSet(testHttpServerEn.Name) {
		testHTTPServerEnabled = ctx.GlobalBool(testHttpServerEn.Name)
	}

	if testHTTPServerEnabled {
		log.Info("Starting test HTTP server handling the requests...")
		testServer = testing.NewTestHttpServer()
		log.Info("Test HTTP server running at " + testServer.URL())

		testCfg := &config.Config{
			GeneralSettings: config.GeneralSettingsConfig{
				RequestTimeoutSec:                        10,
				HeartbeatCacheValidityDurationSec:        60,
				ValStatsCacheValidityDurationSec:         60,
				EconomicsMetricsCacheValidityDurationSec: 6,
				FaucetValue:                              "10000000000",
			},
			ApiLogging: config.ApiLoggingConfig{
				LoggingEnabled:          true,
				ThresholdInMicroSeconds: 10000,
			},
			Observers: []*data.NodeData{
				{
					ShardId: 0,
					Address: testServer.URL(),
				},
				{
					ShardId: 1,
					Address: testServer.URL(),
				},
				{
					ShardId: core.MetachainShardId,
					Address: testServer.URL(),
				},
			},
			FullHistoryNodes: []*data.NodeData{
				{
					ShardId: 0,
					Address: testServer.URL(),
				},
				{
					ShardId: 1,
					Address: testServer.URL(),
				},
				{
					ShardId: core.MetachainShardId,
					Address: testServer.URL(),
				},
			},
			AddressPubkeyConverter: cfg.AddressPubkeyConverter,
			Marshalizer:            erdConfig.TypeConfig{Type: "json"},
			Hasher:                 erdConfig.TypeConfig{Type: "sha256"},
		}

		return createVersionsRegistry(
			testCfg,
			configurationFilePath,
			exCfg,
			statusMetricsHandler,
			ctx.GlobalString(walletKeyPemFile.Name),
			ctx.GlobalString(apiConfigDirectory.Name),
			closableComponents,
			false,
		)
	}

	isRosettaModeEnabled := ctx.GlobalBool(startAsRosetta.Name)
	return createVersionsRegistry(
		cfg,
		configurationFilePath,
		exCfg,
		statusMetricsHandler,
		ctx.GlobalString(walletKeyPemFile.Name),
		ctx.GlobalString(apiConfigDirectory.Name),
		closableComponents,
		isRosettaModeEnabled,
	)
}

func createVersionsRegistry(
	cfg *config.Config,
	configurationFilePath string,
	exCfg *erdConfig.ExternalConfig,
	statusMetricsHandler data.StatusMetricsProvider,
	pemFileLocation string,
	apiConfigDirectoryPath string,
	closableComponents *data.ClosableComponentsHandler,
	isRosettaModeEnabled bool,
) (data.VersionsRegistryHandler, error) {
	pubKeyConverter, err := factory.NewPubkeyConverter(cfg.AddressPubkeyConverter)
	if err != nil {
		return nil, err
	}

	marshalizer, err := marshalFactory.NewMarshalizer(cfg.Marshalizer.Type)
	if err != nil {
		return nil, err
	}
	hasher, err := hasherFactory.NewHasher(cfg.Hasher.Type)
	if err != nil {
		return nil, err
	}

	shardCoord, err := getShardCoordinator(cfg)
	if err != nil {
		return nil, err
	}

	nodesProviderFactory, err := observer.NewNodesProviderFactory(*cfg, configurationFilePath)
	if err != nil {
		return nil, err
	}

	observersProvider, err := nodesProviderFactory.CreateObservers()
	if err != nil {
		return nil, err
	}

	fullHistoryNodesProvider, err := nodesProviderFactory.CreateFullHistoryNodes()
	if err != nil {
		if err != observer.ErrEmptyObserversList {
			return nil, err
		}
	}

	bp, err := process.NewBaseProcessor(
		cfg.GeneralSettings.RequestTimeoutSec,
		shardCoord,
		observersProvider,
		fullHistoryNodesProvider,
		pubKeyConverter,
	)
	if err != nil {
		return nil, err
	}
	bp.StartNodesSyncStateChecks()

	connector, err := createElasticSearchConnector(exCfg)
	if err != nil {
		return nil, err
	}

	accntProc, err := process.NewAccountProcessor(bp, pubKeyConverter, connector)
	if err != nil {
		return nil, err
	}

	faucetValue := big.NewInt(0)
	faucetValue.SetString(cfg.GeneralSettings.FaucetValue, 10)
	faucetProc, err := processFactory.CreateFaucetProcessor(bp, shardCoord, faucetValue, pubKeyConverter, pemFileLocation)
	if err != nil {
		return nil, err
	}

	txProc, err := processFactory.CreateTransactionProcessor(
		bp,
		pubKeyConverter,
		hasher,
		marshalizer,
	)
	if err != nil {
		return nil, err
	}

	scQueryProc, err := process.NewSCQueryProcessor(bp, pubKeyConverter)
	if err != nil {
		return nil, err
	}

	htbCacher := cache.NewHeartbeatMemoryCacher()
	cacheValidity := time.Duration(cfg.GeneralSettings.HeartbeatCacheValidityDurationSec) * time.Second

	htbProc, err := process.NewHeartbeatProcessor(bp, htbCacher, cacheValidity)
	if err != nil {
		return nil, err
	}

	valStatsCacher := cache.NewValidatorsStatsMemoryCacher()
	cacheValidity = time.Duration(cfg.GeneralSettings.ValStatsCacheValidityDurationSec) * time.Second

	valStatsProc, err := process.NewValidatorStatisticsProcessor(bp, valStatsCacher, cacheValidity)
	if err != nil {
		return nil, err
	}

	economicMetricsCacher := cache.NewGenericApiResponseMemoryCacher()
	cacheValidity = time.Duration(cfg.GeneralSettings.EconomicsMetricsCacheValidityDurationSec) * time.Second

	nodeStatusProc, err := process.NewNodeStatusProcessor(bp, economicMetricsCacher, cacheValidity)
	if err != nil {
		return nil, err
	}

	closableComponents.Add(htbProc, valStatsProc, nodeStatusProc, bp)
	if !isRosettaModeEnabled {
		htbProc.StartCacheUpdate()
		valStatsProc.StartCacheUpdate()
		nodeStatusProc.StartCacheUpdate()
	}

	blockProc, err := process.NewBlockProcessor(connector, bp)
	if err != nil {
		return nil, err
	}

	blocksPrc, err := process.NewBlocksProcessor(bp)
	if err != nil {
		return nil, err
	}

	proofProc, err := process.NewProofProcessor(bp, pubKeyConverter)
	if err != nil {
		return nil, err
	}

	esdtSuppliesProc, err := process.NewESDTSupplyProcessor(bp, scQueryProc)
	if err != nil {
		return nil, err
	}

	statusProc, err := process.NewStatusProcessor(bp, statusMetricsHandler)
	if err != nil {
		return nil, err
	}

	facadeArgs := versionsFactory.FacadeArgs{
		ActionsProcessor:             bp,
		AccountProcessor:             accntProc,
		FaucetProcessor:              faucetProc,
		BlockProcessor:               blockProc,
		BlocksProcessor:              blocksPrc,
		HeartbeatProcessor:           htbProc,
		NodeStatusProcessor:          nodeStatusProc,
		ScQueryProcessor:             scQueryProc,
		TransactionProcessor:         txProc,
		ValidatorStatisticsProcessor: valStatsProc,
		ProofProcessor:               proofProc,
		PubKeyConverter:              pubKeyConverter,
		ESDTSuppliesProcessor:        esdtSuppliesProc,
		StatusProcessor:              statusProc,
	}

	apiConfigParser, err := versionsFactory.NewApiConfigParser(apiConfigDirectoryPath)
	if err != nil {
		return nil, err
	}

	return versionsFactory.CreateVersionsRegistry(facadeArgs, apiConfigParser)
}

func createElasticSearchConnector(exCfg *erdConfig.ExternalConfig) (process.ExternalStorageConnector, error) {
	if !exCfg.ElasticSearchConnector.Enabled {
		return database.NewDisabledElasticSearchConnector(), nil
	}

	return database.NewElasticSearchConnector(
		exCfg.ElasticSearchConnector.URL,
		exCfg.ElasticSearchConnector.Username,
		exCfg.ElasticSearchConnector.Password,
	)
}

func getShardCoordinator(cfg *config.Config) (sharding.Coordinator, error) {
	maxShardID := uint32(0)
	for _, obs := range cfg.Observers {
		shardID := obs.ShardId
		isMetaChain := shardID == core.MetachainShardId
		if maxShardID < shardID && !isMetaChain {
			maxShardID = shardID
		}
	}

	shardCoordinator, err := sharding.NewMultiShardCoordinator(maxShardID+1, 0)
	if err != nil {
		return nil, err
	}

	return shardCoordinator, nil
}

func startWebServer(
	versionsRegistry data.VersionsRegistryHandler,
	cliContext *cli.Context,
	generalConfig *config.Config,
	credentialsConfig config.CredentialsConfig,
	statusMetricsProvider data.StatusMetricsProvider,
	isProfileModeActivated bool,
) (*http.Server, error) {
	var err error
	var httpServer *http.Server

	port := generalConfig.GeneralSettings.ServerPort
	asRosetta := cliContext.GlobalBool(startAsRosetta.Name)
	if asRosetta {
		isRosettaOffline := cliContext.GlobalBool(rosettaOffline.Name)
		offlineConfigPath := cliContext.GlobalString(rosettaOfflineConfig.Name)
		var facades map[string]*data.VersionData
		facades, err = versionsRegistry.GetAllVersions()
		if err != nil {
			return nil, err
		}
		httpServer, err = rosetta.CreateServer(facades["v1.0"].Facade, generalConfig, port, isRosettaOffline, offlineConfigPath)
	} else {
		if generalConfig.GeneralSettings.RateLimitWindowDurationSeconds <= 0 {
			return nil, fmt.Errorf("invalid value %d for RateLimitWindowDurationSeconds. It must be greater "+
				"than zero", generalConfig.GeneralSettings.RateLimitWindowDurationSeconds)
		}
		httpServer, err = api.CreateServer(
			versionsRegistry,
			port,
			generalConfig.ApiLogging,
			credentialsConfig,
			statusMetricsProvider,
			generalConfig.GeneralSettings.RateLimitWindowDurationSeconds,
			isProfileModeActivated,
		)
	}
	if err != nil {
		return nil, err
	}
	go func() {
		err = httpServer.ListenAndServe()
		if err != nil {
			log.Error("cannot ListenAndServe()", "err", err)
			os.Exit(1)
		}
	}()

	return httpServer, nil
}

func waitForServerShutdown(httpServer *http.Server, closableComponents *data.ClosableComponentsHandler) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	closableComponents.Close()

	shutdownContext, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_ = httpServer.Shutdown(shutdownContext)
	_ = httpServer.Close()
}

func removeLogColors() {
	err := logger.RemoveLogObserver(os.Stdout)
	if err != nil {
		panic("error removing log observer: " + err.Error())
	}

	err = logger.AddLogObserver(os.Stdout, &logger.PlainFormatter{})
	if err != nil {
		panic("error setting log observer: " + err.Error())
	}
}

func getWorkingDir(ctx *cli.Context, log logger.Logger) string {
	var workingDir string
	var err error
	if ctx.IsSet(workingDirectory.Name) {
		workingDir = ctx.GlobalString(workingDirectory.Name)
	} else {
		workingDir, err = os.Getwd()
		if err != nil {
			log.LogIfError(err)
			workingDir = ""
		}
	}
	log.Trace("working directory", "path", workingDir)

	return workingDir
}

func loadCredentialsConfig(filepath string) (*config.CredentialsConfig, error) {
	cfg := &config.CredentialsConfig{}
	err := core.LoadTomlFile(cfg, filepath)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
