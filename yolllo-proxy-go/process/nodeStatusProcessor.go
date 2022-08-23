package process

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go/common"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
)

const (
	// NetworkStatusPath represents the path where an observer exposes his network metrics
	NetworkStatusPath = "/network/status"

	// NetworkConfigPath represents the path where an observer exposes his network metrics
	NetworkConfigPath = "/network/config"

	// NodeStatusPath represents the path where an observer exposes his node status metrics
	NodeStatusPath = "/node/status"

	// AllIssuedESDTsPath represents the path where an observer exposes all the issued ESDTs
	AllIssuedESDTsPath = "/network/esdts"

	// NetworkEsdtTokensPrefix represents the prefix for the path where an observer exposes ESDT tokens of a kind
	NetworkEsdtTokensPrefix = "/network/esdt"

	// DelegatedInfoPath represents the path where an observer exposes his network delegated info
	DelegatedInfoPath = "/network/delegated-info"

	// DirectStakedPath represents the path where an observer exposes his network direct staked info
	DirectStakedPath = "/network/direct-staked-info"

	// RatingsConfigPath represents the path where an observer exposes his ratings metrics
	RatingsConfigPath = "/network/ratings"

	// GenesisNodesConfigPath represents the path where an observer exposes genesis nodes config
	GenesisNodesConfigPath = "/network/genesis-nodes"

	// EnableEpochsPath represents the path where an observer exposes all the activation epochs
	EnableEpochsPath = "/network/enable-epochs"
)

// NodeStatusProcessor handles the action needed for fetching data related to status metrics from nodes
type NodeStatusProcessor struct {
	proc                  Processor
	economicMetricsCacher GenericApiResponseCacheHandler
	cacheValidityDuration time.Duration
	cancelFunc            func()
}

// NewNodeStatusProcessor creates a new instance of NodeStatusProcessor
func NewNodeStatusProcessor(
	processor Processor,
	economicMetricsCacher GenericApiResponseCacheHandler,
	cacheValidityDuration time.Duration,
) (*NodeStatusProcessor, error) {
	if check.IfNil(processor) {
		return nil, ErrNilCoreProcessor
	}
	if check.IfNil(economicMetricsCacher) {
		return nil, ErrNilEconomicMetricsCacher
	}
	if cacheValidityDuration <= 0 {
		return nil, ErrInvalidCacheValidityDuration
	}

	return &NodeStatusProcessor{
		proc:                  processor,
		economicMetricsCacher: economicMetricsCacher,
		cacheValidityDuration: cacheValidityDuration,
	}, nil
}

// GetNetworkStatusMetrics will simply forward the network status metrics from an observer in the given shard
func (nsp *NodeStatusProcessor) GetNetworkStatusMetrics(shardID uint32) (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetObservers(shardID)
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseNetworkMetrics *data.GenericAPIResponse

		_, err := nsp.proc.CallGetRestEndPoint(observer.Address, NetworkStatusPath, &responseNetworkMetrics)
		if err != nil {
			log.Error("network metrics request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("network metrics request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseNetworkMetrics, nil

	}

	return nil, ErrSendingRequest
}

// GetNetworkConfigMetrics will simply forward the network config metrics from an observer in the given shard
func (nsp *NodeStatusProcessor) GetNetworkConfigMetrics() (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetAllObservers()
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseNetworkMetrics *data.GenericAPIResponse

		_, err = nsp.proc.CallGetRestEndPoint(observer.Address, NetworkConfigPath, &responseNetworkMetrics)
		if err != nil {
			log.Error("network metrics request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("network metrics request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseNetworkMetrics, nil

	}

	return nil, ErrSendingRequest
}

// GetEnableEpochsMetrics will simply forward the activation epochs config metrics from an observer
func (nsp *NodeStatusProcessor) GetEnableEpochsMetrics() (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetAllObservers()
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseEnableEpochsMetrics *data.GenericAPIResponse

		_, err := nsp.proc.CallGetRestEndPoint(observer.Address, EnableEpochsPath, &responseEnableEpochsMetrics)
		if err != nil {
			log.Error("enable epochs metrics request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("enable epochs metrics request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseEnableEpochsMetrics, nil
	}

	return nil, ErrSendingRequest
}

// GetAllIssuedESDTs will forward the issued ESDTs based on the provided type
func (nsp *NodeStatusProcessor) GetAllIssuedESDTs(tokenType string) (*data.GenericAPIResponse, error) {
	if !data.IsValidEsdtPath(tokenType) && tokenType != "" {
		return nil, ErrInvalidTokenType
	}

	observers, err := nsp.proc.GetObservers(core.MetachainShardId)
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseAllIssuedESDTs *data.GenericAPIResponse

		path := AllIssuedESDTsPath
		if tokenType != "" {
			path = fmt.Sprintf("%s/%s", NetworkEsdtTokensPrefix, tokenType)
		}
		_, err := nsp.proc.CallGetRestEndPoint(observer.Address, path, &responseAllIssuedESDTs)
		if err != nil {
			log.Error("all issued esdts request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("all issued esdts request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseAllIssuedESDTs, nil

	}

	return nil, ErrSendingRequest
}

// GetDelegatedInfo returns the delegated info from nodes
func (nsp *NodeStatusProcessor) GetDelegatedInfo() (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetObservers(core.MetachainShardId)
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var delegatedInfoResponse *data.GenericAPIResponse

		_, err := nsp.proc.CallGetRestEndPoint(observer.Address, DelegatedInfoPath, &delegatedInfoResponse)
		if err != nil {
			log.Error("network delegated info request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("network delegated info request", "shard ID", observer.ShardId, "observer", observer.Address)
		return delegatedInfoResponse, nil

	}

	return nil, ErrSendingRequest
}

// GetDirectStakedInfo returns the delegated info from nodes
func (nsp *NodeStatusProcessor) GetDirectStakedInfo() (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetObservers(core.MetachainShardId)
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var directStakedResponse *data.GenericAPIResponse

		_, err := nsp.proc.CallGetRestEndPoint(observer.Address, DirectStakedPath, &directStakedResponse)
		if err != nil {
			log.Error("network direct staked request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("network direct staked request", "shard ID", observer.ShardId, "observer", observer.Address)
		return directStakedResponse, nil

	}

	return nil, ErrSendingRequest
}

// GetRatingsConfig will simply forward the ratings configuration from an observer
func (nsp *NodeStatusProcessor) GetRatingsConfig() (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetAllObservers()
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseRatingsConfig *data.GenericAPIResponse

		_, err = nsp.proc.CallGetRestEndPoint(observer.Address, RatingsConfigPath, &responseRatingsConfig)
		if err != nil {
			log.Error("ratings metrics request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("ratings metrics request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseRatingsConfig, nil

	}

	return nil, ErrSendingRequest
}

func (nsp *NodeStatusProcessor) getNodeStatusMetrics(shardID uint32) (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetObservers(shardID)
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseNetworkMetrics *data.GenericAPIResponse

		_, err = nsp.proc.CallGetRestEndPoint(observer.Address, NodeStatusPath, &responseNetworkMetrics)
		if err != nil {
			log.Error("node status metrics request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("node status metrics request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseNetworkMetrics, nil

	}

	return nil, ErrSendingRequest
}

// GetLatestFullySynchronizedHyperblockNonce will compute nonce of the latest hyperblock that can be returned
func (nsp *NodeStatusProcessor) GetLatestFullySynchronizedHyperblockNonce() (uint64, error) {
	shardsIDs, err := nsp.getShardsIDs()
	if err != nil {
		return 0, err
	}

	nonces := make([]uint64, 0)
	for shardID := range shardsIDs {
		nodeStatusResponse, err := nsp.getNodeStatusMetrics(shardID)
		if err != nil {
			return 0, err
		}

		if nodeStatusResponse.Error != "" {
			return 0, errors.New(nodeStatusResponse.Error)
		}

		var nonce uint64
		var ok bool
		if shardID == core.MetachainShardId {
			nonce, ok = getNonceFromMetachainStatus(nodeStatusResponse.Data)
		} else {
			nonce, ok = getNonceFromShardStatus(nodeStatusResponse.Data)
		}
		if !ok {
			return 0, ErrCannotParseNodeStatusMetrics
		}

		nonces = append(nonces, nonce)
	}

	return getMinNonce(nonces), nil
}

func getMinNonce(noncesSlice []uint64) uint64 {
	// initialize min with max uint64 value
	min := uint64(math.MaxUint64)
	for _, value := range noncesSlice {
		if value < min {
			min = value
		}
	}

	return min
}

func (nsp *NodeStatusProcessor) getShardsIDs() (map[uint32]struct{}, error) {
	observers, err := nsp.proc.GetAllObservers()
	if err != nil {
		return nil, err
	}

	shardsIDs := make(map[uint32]struct{})
	for _, observer := range observers {
		shardsIDs[observer.ShardId] = struct{}{}
	}

	if len(shardsIDs) == 0 {
		return nil, ErrMissingObserver
	}

	return shardsIDs, nil
}

func getNonceFromShardStatus(nodeStatusData interface{}) (uint64, bool) {
	metric, ok := getMetric(nodeStatusData, common.MetricCrossCheckBlockHeight)
	if !ok {
		return 0, false
	}

	return parseMetricCrossCheckBlockHeight(metric)
}

func getNonceFromMetachainStatus(nodeStatusData interface{}) (uint64, bool) {
	metric, ok := getMetric(nodeStatusData, common.MetricNonce)
	if !ok {
		return 0, false
	}

	return getUint(metric), true
}

func getMetric(nodeStatusData interface{}, metric string) (interface{}, bool) {
	metricsMapI, ok := nodeStatusData.(map[string]interface{})
	if !ok {
		return nil, false
	}

	metricsMap, ok := metricsMapI["metrics"]
	if !ok {
		return nil, false
	}

	metrics, ok := metricsMap.(map[string]interface{})
	if !ok {
		return nil, false
	}

	value, ok := metrics[metric]
	if !ok {
		return nil, false
	}

	return value, true
}

func parseMetricCrossCheckBlockHeight(value interface{}) (uint64, bool) {
	valueStr, ok := value.(string)
	if !ok {
		return 0, false
	}

	// metric looks like that
	// "meta 886717"
	values := strings.Split(valueStr, " ")
	if len(values) < 2 {
		return 0, false
	}

	nonce, err := strconv.ParseUint(values[1], 10, 64)
	if err != nil {
		return 0, false
	}

	return nonce, true
}

func getUint(value interface{}) uint64 {
	valueFloat, ok := value.(float64)
	if !ok {
		return 0
	}

	return uint64(valueFloat)
}

// GetGenesisNodesPubKeys will return genesis nodes public keys
func (nsp *NodeStatusProcessor) GetGenesisNodesPubKeys() (*data.GenericAPIResponse, error) {
	observers, err := nsp.proc.GetAllObservers()
	if err != nil {
		return nil, err
	}

	for _, observer := range observers {
		var responseGenesisNodesConfig *data.GenericAPIResponse

		_, err = nsp.proc.CallGetRestEndPoint(observer.Address, GenesisNodesConfigPath, &responseGenesisNodesConfig)
		if err != nil {
			log.Error("genesis nodes request", "observer", observer.Address, "error", err.Error())
			continue
		}

		log.Info("genesis nodes request", "shard ID", observer.ShardId, "observer", observer.Address)
		return responseGenesisNodesConfig, nil

	}

	return nil, ErrSendingRequest
}
