package facade

import (
	"encoding/json"
	"math/big"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	"github.com/ElrondNetwork/elrond-proxy-go/api/groups"
	"github.com/ElrondNetwork/elrond-proxy-go/common"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
)

// interfaces assertions. verifies that all API endpoint have their corresponding methods in the facade
var _ groups.ActionsFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.AccountsFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.BlockFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.BlocksFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.BlockAtlasFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.HyperBlockFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.NetworkFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.NodeFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.TransactionFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.ValidatorFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.VmValuesFacadeHandler = (*ElrondProxyFacade)(nil)
var _ groups.ProofFacadeHandler = (*ElrondProxyFacade)(nil)

// ElrondProxyFacade implements the facade used in api calls
type ElrondProxyFacade struct {
	actionsProc      ActionsProcessor
	accountProc      AccountProcessor
	txProc           TransactionProcessor
	scQueryService   SCQueryService
	heartbeatProc    HeartbeatProcessor
	valStatsProc     ValidatorStatisticsProcessor
	faucetProc       FaucetProcessor
	nodeStatusProc   NodeStatusProcessor
	blockProc        BlockProcessor
	blocksProc       BlocksProcessor
	proofProc        ProofProcessor
	esdtSuppliesProc ESDTSupplyProcessor
	statusProc       StatusProcessor

	pubKeyConverter core.PubkeyConverter
}

// NewElrondProxyFacade creates a new ElrondProxyFacade instance
func NewElrondProxyFacade(
	actionsProc ActionsProcessor,
	accountProc AccountProcessor,
	txProc TransactionProcessor,
	scQueryService SCQueryService,
	heartbeatProc HeartbeatProcessor,
	valStatsProc ValidatorStatisticsProcessor,
	faucetProc FaucetProcessor,
	nodeStatusProc NodeStatusProcessor,
	blockProc BlockProcessor,
	blocksProc BlocksProcessor,
	proofProc ProofProcessor,
	pubKeyConverter core.PubkeyConverter,
	esdtSuppliesProc ESDTSupplyProcessor,
	statusProc StatusProcessor,
) (*ElrondProxyFacade, error) {
	if actionsProc == nil {
		return nil, ErrNilActionsProcessor
	}
	if accountProc == nil {
		return nil, ErrNilAccountProcessor
	}
	if txProc == nil {
		return nil, ErrNilTransactionProcessor
	}
	if scQueryService == nil {
		return nil, ErrNilSCQueryService
	}
	if heartbeatProc == nil {
		return nil, ErrNilHeartbeatProcessor
	}
	if valStatsProc == nil {
		return nil, ErrNilValidatorStatisticsProcessor
	}
	if faucetProc == nil {
		return nil, ErrNilFaucetProcessor
	}
	if nodeStatusProc == nil {
		return nil, ErrNilNodeStatusProcessor
	}
	if blockProc == nil {
		return nil, ErrNilBlockProcessor
	}
	if blocksProc == nil {
		return nil, ErrNilBlocksProcessor
	}
	if proofProc == nil {
		return nil, ErrNilProofProcessor
	}
	if esdtSuppliesProc == nil {
		return nil, ErrNilESDTSuppliesProcessor
	}
	if statusProc == nil {
		return nil, ErrNilStatusProcessor
	}

	return &ElrondProxyFacade{
		actionsProc:      actionsProc,
		accountProc:      accountProc,
		txProc:           txProc,
		scQueryService:   scQueryService,
		heartbeatProc:    heartbeatProc,
		valStatsProc:     valStatsProc,
		faucetProc:       faucetProc,
		nodeStatusProc:   nodeStatusProc,
		blockProc:        blockProc,
		blocksProc:       blocksProc,
		proofProc:        proofProc,
		pubKeyConverter:  pubKeyConverter,
		esdtSuppliesProc: esdtSuppliesProc,
		statusProc:       statusProc,
	}, nil
}

// GetAccount returns an account based on the input address
func (epf *ElrondProxyFacade) GetAccount(address string) (*data.Account, error) {
	return epf.accountProc.GetAccount(address)
}

// GetKeyValuePairs returns the key-value pairs for the given address
func (epf *ElrondProxyFacade) GetKeyValuePairs(address string) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetKeyValuePairs(address)
}

// GetValueForKey returns the value for the given address and key
func (epf *ElrondProxyFacade) GetValueForKey(address string, key string) (string, error) {
	return epf.accountProc.GetValueForKey(address, key)
}

// GetShardIDForAddress returns the computed shard ID for the given address based on the current proxy's configuration
func (epf *ElrondProxyFacade) GetShardIDForAddress(address string) (uint32, error) {
	return epf.accountProc.GetShardIDForAddress(address)
}

// GetTransactions returns transactions by address
func (epf *ElrondProxyFacade) GetTransactions(address string) ([]data.DatabaseTransaction, error) {
	return epf.accountProc.GetTransactions(address)
}

// GetESDTTokenData returns the token data for a given token name
func (epf *ElrondProxyFacade) GetESDTTokenData(address string, key string) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetESDTTokenData(address, key)
}

// GetESDTNftTokenData returns the token data for a given token name
func (epf *ElrondProxyFacade) GetESDTNftTokenData(address string, key string, nonce uint64) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetESDTNftTokenData(address, key, nonce)
}

// GetESDTsWithRole returns the tokens where the given address has the assigned role
func (epf *ElrondProxyFacade) GetESDTsWithRole(address string, role string) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetESDTsWithRole(address, role)
}

// GetESDTsRoles returns the tokens and roles for the given address
func (epf *ElrondProxyFacade) GetESDTsRoles(address string) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetESDTsRoles(address)
}

// GetNFTTokenIDsRegisteredByAddress returns the token identifiers of the NFTs registered by the address
func (epf *ElrondProxyFacade) GetNFTTokenIDsRegisteredByAddress(address string) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetNFTTokenIDsRegisteredByAddress(address)
}

// GetAllESDTTokens returns all the ESDT tokens for a given address
func (epf *ElrondProxyFacade) GetAllESDTTokens(address string) (*data.GenericAPIResponse, error) {
	return epf.accountProc.GetAllESDTTokens(address)
}

// SendTransaction should send the transaction to the correct observer
func (epf *ElrondProxyFacade) SendTransaction(tx *data.Transaction) (int, string, error) {
	return epf.txProc.SendTransaction(tx)
}

// SendMultipleTransactions should send the transactions to the correct observers
func (epf *ElrondProxyFacade) SendMultipleTransactions(txs []*data.Transaction) (data.MultipleTransactionsResponseData, error) {
	return epf.txProc.SendMultipleTransactions(txs)
}

// SimulateTransaction should send the transaction to the correct observer for simulation
func (epf *ElrondProxyFacade) SimulateTransaction(tx *data.Transaction, checkSignature bool) (*data.GenericAPIResponse, error) {
	return epf.txProc.SimulateTransaction(tx, checkSignature)
}

// TransactionCostRequest should return how many gas units a transaction will cost
func (epf *ElrondProxyFacade) TransactionCostRequest(tx *data.Transaction) (*data.TxCostResponseData, error) {
	return epf.txProc.TransactionCostRequest(tx)
}

// GetTransactionStatus should return transaction status
func (epf *ElrondProxyFacade) GetTransactionStatus(txHash string, sender string) (string, error) {
	return epf.txProc.GetTransactionStatus(txHash, sender)
}

// GetTransaction should return a transaction by hash
func (epf *ElrondProxyFacade) GetTransaction(txHash string, withResults bool) (*data.FullTransaction, error) {
	return epf.txProc.GetTransaction(txHash, withResults)
}

// ReloadObservers will try to reload the observers
func (epf *ElrondProxyFacade) ReloadObservers() data.NodesReloadResponse {
	return epf.actionsProc.ReloadObservers()
}

// ReloadFullHistoryObservers will try to reload the full history observers
func (epf *ElrondProxyFacade) ReloadFullHistoryObservers() data.NodesReloadResponse {
	return epf.actionsProc.ReloadFullHistoryObservers()
}

// GetTransactionByHashAndSenderAddress should return a transaction by hash and sender address
func (epf *ElrondProxyFacade) GetTransactionByHashAndSenderAddress(txHash string, sndAddr string, withEvents bool) (*data.FullTransaction, int, error) {
	return epf.txProc.GetTransactionByHashAndSenderAddress(txHash, sndAddr, withEvents)
}

// IsFaucetEnabled returns true if the faucet mechanism is enabled or false otherwise
func (epf *ElrondProxyFacade) IsFaucetEnabled() bool {
	return epf.faucetProc.IsEnabled()
}

// SendUserFunds should send a transaction to load one user's account with extra funds from an account in the pem file
func (epf *ElrondProxyFacade) SendUserFunds(receiver string, value *big.Int) error {
	senderSk, senderPk, err := epf.faucetProc.SenderDetailsFromPem(receiver)
	if err != nil {
		return err
	}

	senderAccount, err := epf.accountProc.GetAccount(senderPk)
	if err != nil {
		return err
	}

	networkCfg, err := epf.getNetworkConfig()
	if err != nil {
		return err
	}

	tx, err := epf.faucetProc.GenerateTxForSendUserFunds(
		senderSk,
		senderPk,
		senderAccount.Nonce,
		receiver,
		value,
		networkCfg,
	)
	if err != nil {
		return err
	}

	_, _, err = epf.txProc.SendTransaction(tx)
	return err
}

func (epf *ElrondProxyFacade) getNetworkConfig() (*data.NetworkConfig, error) {
	genericResponse, err := epf.nodeStatusProc.GetNetworkConfigMetrics()
	if err != nil {
		return nil, err
	}

	networkConfigBytes, err := json.Marshal(&genericResponse.Data)
	if err != nil {
		return nil, err
	}

	networkCfg := &data.NetworkConfig{}
	err = json.Unmarshal(networkConfigBytes, networkCfg)

	return networkCfg, err
}

// ExecuteSCQuery retrieves data from existing SC trie through the use of a VM
func (epf *ElrondProxyFacade) ExecuteSCQuery(query *data.SCQuery) (*vm.VMOutputApi, error) {
	return epf.scQueryService.ExecuteQuery(query)
}

// GetHeartbeatData retrieves the heartbeat status from one observer
func (epf *ElrondProxyFacade) GetHeartbeatData() (*data.HeartbeatResponse, error) {
	return epf.heartbeatProc.GetHeartbeatData()
}

// GetNetworkConfigMetrics retrieves the node's configuration's metrics
func (epf *ElrondProxyFacade) GetNetworkConfigMetrics() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetNetworkConfigMetrics()
}

// GetNetworkStatusMetrics retrieves the node's network metrics for a given shard
func (epf *ElrondProxyFacade) GetNetworkStatusMetrics(shardID uint32) (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetNetworkStatusMetrics(shardID)
}

// GetESDTSupply retrieves the supply for the provided token
func (epf *ElrondProxyFacade) GetESDTSupply(token string) (*data.ESDTSupplyResponse, error) {
	return epf.esdtSuppliesProc.GetESDTSupply(token)
}

// GetEconomicsDataMetrics retrieves the node's network metrics for a given shard
func (epf *ElrondProxyFacade) GetEconomicsDataMetrics() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetEconomicsDataMetrics()
}

// GetDelegatedInfo retrieves the node's network delegated info
func (epf *ElrondProxyFacade) GetDelegatedInfo() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetDelegatedInfo()
}

// GetDirectStakedInfo retrieves the node's direct staked values
func (epf *ElrondProxyFacade) GetDirectStakedInfo() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetDirectStakedInfo()
}

// GetAllIssuedESDTs retrieves all the issued ESDTs from the node
func (epf *ElrondProxyFacade) GetAllIssuedESDTs(tokenType string) (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetAllIssuedESDTs(tokenType)
}

// GetEnableEpochsMetrics retrieves the activation epochs
func (epf *ElrondProxyFacade) GetEnableEpochsMetrics() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetEnableEpochsMetrics()
}

// GetRatingsConfig retrieves the node's configuration's metrics
func (epf *ElrondProxyFacade) GetRatingsConfig() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetRatingsConfig()
}

// GetBlockByHash retrieves the block by hash for a given shard
func (epf *ElrondProxyFacade) GetBlockByHash(shardID uint32, hash string, withTxs bool) (*data.BlockApiResponse, error) {
	return epf.blockProc.GetBlockByHash(shardID, hash, withTxs)
}

// GetBlockByNonce retrieves the block by nonce for a given shard
func (epf *ElrondProxyFacade) GetBlockByNonce(shardID uint32, nonce uint64, withTxs bool) (*data.BlockApiResponse, error) {
	return epf.blockProc.GetBlockByNonce(shardID, nonce, withTxs)
}

// GetBlocksByRound retrieves the blocks for a given round
func (epf *ElrondProxyFacade) GetBlocksByRound(round uint64, withTxs bool) (*data.BlocksApiResponse, error) {
	return epf.blocksProc.GetBlocksByRound(round, withTxs)
}

// GetInternalBlockByHash retrieves the internal block by hash for a given shard
func (epf *ElrondProxyFacade) GetInternalBlockByHash(shardID uint32, hash string, format common.OutputFormat) (*data.InternalBlockApiResponse, error) {
	return epf.blockProc.GetInternalBlockByHash(shardID, hash, format)
}

// GetInternalBlockByNonce retrieves the internal block by nonce for a given shard
func (epf *ElrondProxyFacade) GetInternalBlockByNonce(shardID uint32, nonce uint64, format common.OutputFormat) (*data.InternalBlockApiResponse, error) {
	return epf.blockProc.GetInternalBlockByNonce(shardID, nonce, format)
}

// GetInternalStartOfEpochMetaBlock retrieves the internal block by nonce for a given shard
func (epf *ElrondProxyFacade) GetInternalStartOfEpochMetaBlock(epoch uint32, format common.OutputFormat) (*data.InternalBlockApiResponse, error) {
	return epf.blockProc.GetInternalStartOfEpochMetaBlock(epoch, format)
}

// GetInternalMiniBlockByHash retrieves the internal miniblock by hash for a given shard
func (epf *ElrondProxyFacade) GetInternalMiniBlockByHash(shardID uint32, hash string, epoch uint32, format common.OutputFormat) (*data.InternalMiniBlockApiResponse, error) {
	return epf.blockProc.GetInternalMiniBlockByHash(shardID, hash, epoch, format)
}

// GetHyperBlockByHash retrieves the hyperblock by hash
func (epf *ElrondProxyFacade) GetHyperBlockByHash(hash string) (*data.HyperblockApiResponse, error) {
	return epf.blockProc.GetHyperBlockByHash(hash)
}

// GetHyperBlockByNonce retrieves the block by nonce
func (epf *ElrondProxyFacade) GetHyperBlockByNonce(nonce uint64) (*data.HyperblockApiResponse, error) {
	return epf.blockProc.GetHyperBlockByNonce(nonce)
}

// ValidatorStatistics will return the statistics from an observer
func (epf *ElrondProxyFacade) ValidatorStatistics() (map[string]*data.ValidatorApiResponse, error) {
	valStats, err := epf.valStatsProc.GetValidatorStatistics()
	if err != nil {
		return nil, err
	}

	return valStats.Statistics, nil
}

// GetAtlasBlockByShardIDAndNonce returns block by shardID and nonce in a BlockAtlas-friendly-format
func (epf *ElrondProxyFacade) GetAtlasBlockByShardIDAndNonce(shardID uint32, nonce uint64) (data.AtlasBlock, error) {
	return epf.blockProc.GetAtlasBlockByShardIDAndNonce(shardID, nonce)
}

// GetAddressConverter returns the address converter
func (epf *ElrondProxyFacade) GetAddressConverter() (core.PubkeyConverter, error) {
	return epf.pubKeyConverter, nil
}

// GetLatestFullySynchronizedHyperblockNonce returns the latest fully synchronized hyperblock nonce
func (epf *ElrondProxyFacade) GetLatestFullySynchronizedHyperblockNonce() (uint64, error) {
	return epf.nodeStatusProc.GetLatestFullySynchronizedHyperblockNonce()
}

// ComputeTransactionHash will compute hash of a given transaction
func (epf *ElrondProxyFacade) ComputeTransactionHash(tx *data.Transaction) (string, error) {
	return epf.txProc.ComputeTransactionHash(tx)
}

// GetProof returns the Merkle proof for the given address
func (epf *ElrondProxyFacade) GetProof(rootHash string, address string) (*data.GenericAPIResponse, error) {
	return epf.proofProc.GetProof(rootHash, address)
}

// GetProofCurrentRootHash returns the Merkle proof for the given address
func (epf *ElrondProxyFacade) GetProofCurrentRootHash(address string) (*data.GenericAPIResponse, error) {
	return epf.proofProc.GetProofCurrentRootHash(address)
}

// VerifyProof verifies the given Merkle proof
func (epf *ElrondProxyFacade) VerifyProof(rootHash string, address string, proof []string) (*data.GenericAPIResponse, error) {
	return epf.proofProc.VerifyProof(rootHash, address, proof)
}

// GetMetrics will return the status metrics
func (epf *ElrondProxyFacade) GetMetrics() map[string]*data.EndpointMetrics {
	return epf.statusProc.GetMetrics()
}

// GetMetricsForPrometheus will return the status metrics in a prometheus format
func (epf *ElrondProxyFacade) GetMetricsForPrometheus() string {
	return epf.statusProc.GetMetricsForPrometheus()
}

// GetGenesisNodesPubKeys retrieves the node's configuration public keys
func (epf *ElrondProxyFacade) GetGenesisNodesPubKeys() (*data.GenericAPIResponse, error) {
	return epf.nodeStatusProc.GetGenesisNodesPubKeys()
}
