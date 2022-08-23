package groups

import (
	"math/big"

	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	"github.com/ElrondNetwork/elrond-proxy-go/common"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
)

// AccountsFacadeHandler interface defines methods that can be used from the facade
type AccountsFacadeHandler interface {
	GetAccount(address string) (*data.Account, error)
	GetTransactions(address string) ([]data.DatabaseTransaction, error)
	GetShardIDForAddress(address string) (uint32, error)
	GetValueForKey(address string, key string) (string, error)
	GetAllESDTTokens(address string) (*data.GenericAPIResponse, error)
	GetKeyValuePairs(address string) (*data.GenericAPIResponse, error)
	GetESDTTokenData(address string, key string) (*data.GenericAPIResponse, error)
	GetESDTsWithRole(address string, role string) (*data.GenericAPIResponse, error)
	GetESDTsRoles(address string) (*data.GenericAPIResponse, error)
	GetESDTNftTokenData(address string, key string, nonce uint64) (*data.GenericAPIResponse, error)
	GetNFTTokenIDsRegisteredByAddress(address string) (*data.GenericAPIResponse, error)
}

// BlockFacadeHandler interface defines methods that can be used from the facade
type BlockFacadeHandler interface {
	GetBlockByNonce(shardID uint32, nonce uint64, withTxs bool) (*data.BlockApiResponse, error)
	GetBlockByHash(shardID uint32, hash string, withTxs bool) (*data.BlockApiResponse, error)
}

// BlocksFacadeHandler interface defines methods that can be used from the facade
type BlocksFacadeHandler interface {
	GetBlocksByRound(round uint64, withTxs bool) (*data.BlocksApiResponse, error)
}

// InternalFacadeHandler interface defines methods that can be used from facade context variable
type InternalFacadeHandler interface {
	GetInternalBlockByHash(shardID uint32, hash string, format common.OutputFormat) (*data.InternalBlockApiResponse, error)
	GetInternalBlockByNonce(shardID uint32, round uint64, format common.OutputFormat) (*data.InternalBlockApiResponse, error)
	GetInternalMiniBlockByHash(shardID uint32, hash string, epoch uint32, format common.OutputFormat) (*data.InternalMiniBlockApiResponse, error)
	GetInternalStartOfEpochMetaBlock(epoch uint32, format common.OutputFormat) (*data.InternalBlockApiResponse, error)
}

// BlockAtlasFacadeHandler interface defines methods that can be used from facade context variable
type BlockAtlasFacadeHandler interface {
	GetAtlasBlockByShardIDAndNonce(shardID uint32, nonce uint64) (data.AtlasBlock, error)
}

// HyperBlockFacadeHandler defines the actions needed for fetching the hyperblocks from the nodes
type HyperBlockFacadeHandler interface {
	GetHyperBlockByNonce(nonce uint64) (*data.HyperblockApiResponse, error)
	GetHyperBlockByHash(hash string) (*data.HyperblockApiResponse, error)
}

// NetworkFacadeHandler interface defines methods that can be used from the facade
type NetworkFacadeHandler interface {
	GetNetworkStatusMetrics(shardID uint32) (*data.GenericAPIResponse, error)
	GetNetworkConfigMetrics() (*data.GenericAPIResponse, error)
	GetEconomicsDataMetrics() (*data.GenericAPIResponse, error)
	GetAllIssuedESDTs(tokenType string) (*data.GenericAPIResponse, error)
	GetDirectStakedInfo() (*data.GenericAPIResponse, error)
	GetDelegatedInfo() (*data.GenericAPIResponse, error)
	GetEnableEpochsMetrics() (*data.GenericAPIResponse, error)
	GetESDTSupply(token string) (*data.ESDTSupplyResponse, error)
	GetRatingsConfig() (*data.GenericAPIResponse, error)
	GetGenesisNodesPubKeys() (*data.GenericAPIResponse, error)
}

// NodeFacadeHandler interface defines methods that can be used from the facade
type NodeFacadeHandler interface {
	GetHeartbeatData() (*data.HeartbeatResponse, error)
}

// StatusFacadeHandler interface defines methods that can be used from the facade
type StatusFacadeHandler interface {
	GetMetrics() map[string]*data.EndpointMetrics
	GetMetricsForPrometheus() string
}

// TransactionFacadeHandler interface defines methods that can be used from the facade
type TransactionFacadeHandler interface {
	SendTransaction(tx *data.Transaction) (int, string, error)
	SendMultipleTransactions(txs []*data.Transaction) (data.MultipleTransactionsResponseData, error)
	SimulateTransaction(tx *data.Transaction, checkSignature bool) (*data.GenericAPIResponse, error)
	IsFaucetEnabled() bool
	SendUserFunds(receiver string, value *big.Int) error
	TransactionCostRequest(tx *data.Transaction) (*data.TxCostResponseData, error)
	GetTransactionStatus(txHash string, sender string) (string, error)
	GetTransaction(txHash string, withResults bool) (*data.FullTransaction, error)
	GetTransactionByHashAndSenderAddress(txHash string, sndAddr string, withEvents bool) (*data.FullTransaction, int, error)
}

// ProofFacadeHandler interface defines methods that can be used from the facade
type ProofFacadeHandler interface {
	GetProof(rootHash string, address string) (*data.GenericAPIResponse, error)
	GetProofCurrentRootHash(address string) (*data.GenericAPIResponse, error)
	VerifyProof(rootHash string, address string, proof []string) (*data.GenericAPIResponse, error)
}

// ValidatorFacadeHandler interface defines methods that can be used from the facade
type ValidatorFacadeHandler interface {
	ValidatorStatistics() (map[string]*data.ValidatorApiResponse, error)
}

// VmValuesFacadeHandler interface defines methods that can be used from the facade
type VmValuesFacadeHandler interface {
	ExecuteSCQuery(*data.SCQuery) (*vm.VMOutputApi, error)
}

// ActionsFacadeHandler interface defines methods that can be used from the facade
type ActionsFacadeHandler interface {
	ReloadObservers() data.NodesReloadResponse
	ReloadFullHistoryObservers() data.NodesReloadResponse
}
