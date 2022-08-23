package process

import (
	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/data/transaction"
	"github.com/ElrondNetwork/elrond-go-core/data/vm"
	"github.com/ElrondNetwork/elrond-go-crypto"
	"github.com/ElrondNetwork/elrond-go/sharding"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
	"github.com/ElrondNetwork/elrond-proxy-go/observer"
)

// Processor defines what a processor should be able to do
type Processor interface {
	GetObservers(shardID uint32) ([]*data.NodeData, error)
	GetAllObservers() ([]*data.NodeData, error)
	GetObserversOnePerShard() ([]*data.NodeData, error)
	GetFullHistoryNodesOnePerShard() ([]*data.NodeData, error)
	GetFullHistoryNodes(shardID uint32) ([]*data.NodeData, error)
	GetAllFullHistoryNodes() ([]*data.NodeData, error)
	GetShardIDs() []uint32
	ComputeShardId(addressBuff []byte) (uint32, error)
	CallGetRestEndPoint(address string, path string, value interface{}) (int, error)
	CallPostRestEndPoint(address string, path string, data interface{}, response interface{}) (int, error)
	GetShardCoordinator() sharding.Coordinator
	GetPubKeyConverter() core.PubkeyConverter
	GetObserverProvider() observer.NodesProviderHandler
	GetFullHistoryNodesProvider() observer.NodesProviderHandler
	IsInterfaceNil() bool
}

// ExternalStorageConnector defines what a external storage connector should be able to do
type ExternalStorageConnector interface {
	GetTransactionsByAddress(address string) ([]data.DatabaseTransaction, error)
	GetAtlasBlockByShardIDAndNonce(shardID uint32, nonce uint64) (data.AtlasBlock, error)
	IsInterfaceNil() bool
}

// PrivateKeysLoaderHandler defines what a component which handles loading of the private keys file should do
type PrivateKeysLoaderHandler interface {
	PrivateKeysByShard() (map[uint32][]crypto.PrivateKey, error)
}

// HeartbeatCacheHandler will define what a real heartbeat cacher should do
type HeartbeatCacheHandler interface {
	LoadHeartbeats() (*data.HeartbeatResponse, error)
	StoreHeartbeats(hbts *data.HeartbeatResponse) error
	IsInterfaceNil() bool
}

// ValidatorStatisticsCacheHandler will define what a real validator statistics cacher should do
type ValidatorStatisticsCacheHandler interface {
	LoadValStats() (map[string]*data.ValidatorApiResponse, error)
	StoreValStats(valStats map[string]*data.ValidatorApiResponse) error
	IsInterfaceNil() bool
}

// GenericApiResponseCacheHandler will define what a real economic metrics cacher should do
type GenericApiResponseCacheHandler interface {
	Load() (*data.GenericAPIResponse, error)
	Store(response *data.GenericAPIResponse)
	IsInterfaceNil() bool
}

// TransactionCostHandler will define what a real transaction cost handler should do
type TransactionCostHandler interface {
	ResolveCostRequest(tx *data.Transaction) (*data.TxCostResponseData, error)
}

// LogsMergerHandler will define what a real merge logs handler should do
type LogsMergerHandler interface {
	MergeLogEvents(logSource *transaction.ApiLogs, logDestination *transaction.ApiLogs) *transaction.ApiLogs
	IsInterfaceNil() bool
}

// SCQueryService defines how data should be get from a SC account
type SCQueryService interface {
	ExecuteQuery(query *data.SCQuery) (*vm.VMOutputApi, error)
	IsInterfaceNil() bool
}

// StatusMetricsProvider defines what a status metrics provider should do
type StatusMetricsProvider interface {
	GetAll() map[string]*data.EndpointMetrics
	GetMetricsForPrometheus() string
	IsInterfaceNil() bool
}
