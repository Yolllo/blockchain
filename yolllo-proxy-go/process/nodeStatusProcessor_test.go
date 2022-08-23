package process

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-proxy-go/data"
	"github.com/ElrondNetwork/elrond-proxy-go/process/mock"
	"github.com/stretchr/testify/require"
)

func TestNewNodeStatusProcessor_NilBaseProcessor(t *testing.T) {
	t.Parallel()

	nodeStatusProc, err := NewNodeStatusProcessor(nil, &mock.GenericApiResponseCacherMock{}, time.Second)

	require.Equal(t, ErrNilCoreProcessor, err)
	require.Nil(t, nodeStatusProc)
}

func TestNewNodeStatusProcessor_NilCacher(t *testing.T) {
	t.Parallel()

	nodeStatusProc, err := NewNodeStatusProcessor(&mock.ProcessorStub{}, nil, time.Second)

	require.Equal(t, ErrNilEconomicMetricsCacher, err)
	require.Nil(t, nodeStatusProc)
}

func TestNewNodeStatusProcessor_InvalidCacheValidityDuration(t *testing.T) {
	t.Parallel()

	nodeStatusProc, err := NewNodeStatusProcessor(&mock.ProcessorStub{}, &mock.GenericApiResponseCacherMock{}, -1*time.Second)

	require.Equal(t, ErrInvalidCacheValidityDuration, err)
	require.Nil(t, nodeStatusProc)
}

func TestNodeStatusProcessor_GetConfigMetricsGetRestEndPointError(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() ([]*data.NodeData, error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			return 0, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetNetworkConfigMetrics()
	require.Equal(t, ErrSendingRequest, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetConfigMetrics(t *testing.T) {
	t.Parallel()

	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() ([]*data.NodeData, error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			localMap := map[string]interface{}{
				"key": 1,
			}
			genericResp := &data.GenericAPIResponse{Data: localMap}
			genRespBytes, _ := json.Marshal(genericResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	genericResponse, err := nodeStatusProc.GetNetworkConfigMetrics()
	require.Nil(t, err)
	require.NotNil(t, genericResponse)

	map1, ok := genericResponse.Data.(map[string]interface{})
	require.True(t, ok)

	valueFromMap, ok := map1["key"]
	require.True(t, ok)
	require.Equal(t, 1, int(valueFromMap.(float64)))

}

func TestNodeStatusProcessor_GetNetworkMetricsGetObserversFailedShouldErr(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return nil, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetNetworkStatusMetrics(0)
	require.Equal(t, localErr, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetNetworkMetricsGetRestEndPointError(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			return 0, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetNetworkStatusMetrics(0)
	require.Equal(t, ErrSendingRequest, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetNetworkMetrics(t *testing.T) {
	t.Parallel()

	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			localMap := map[string]interface{}{
				"key": 1,
			}
			genericResp := &data.GenericAPIResponse{Data: localMap}
			genRespBytes, _ := json.Marshal(genericResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	genericResponse, err := nodeStatusProc.GetNetworkStatusMetrics(0)
	require.Nil(t, err)
	require.NotNil(t, genericResponse)

	map1, ok := genericResponse.Data.(map[string]interface{})
	require.True(t, ok)

	valueFromMap, ok := map1["key"]
	require.True(t, ok)
	require.Equal(t, 1, int(valueFromMap.(float64)))
}

func TestNodeStatusProcessor_GetLatestBlockNonce(t *testing.T) {
	t.Parallel()

	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
				{Address: "address2", ShardId: core.MetachainShardId},
			}, nil
		},
		GetObserversCalled: func(shardId uint32) ([]*data.NodeData, error) {
			if shardId == 0 {
				return []*data.NodeData{
					{Address: "address1", ShardId: 0},
				}, nil
			} else {
				return []*data.NodeData{
					{Address: "address2", ShardId: core.MetachainShardId},
				}, nil
			}
		},

		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {

			var localMap map[string]interface{}
			if address == "address1" {
				localMap = map[string]interface{}{
					"metrics": map[string]interface{}{
						"erd_cross_check_block_height": "meta 123",
					},
				}
			} else {
				localMap = map[string]interface{}{
					"metrics": map[string]interface{}{
						"erd_nonce": 122,
					},
				}
			}

			genericResp := &data.GenericAPIResponse{Data: localMap}
			genRespBytes, _ := json.Marshal(genericResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	nonce, err := nodeStatusProc.GetLatestFullySynchronizedHyperblockNonce()
	require.NoError(t, err)
	require.Equal(t, uint64(122), nonce)
}

func TestNodeStatusProcessor_GetAllIssuedEDTsGetObserversFailedShouldErr(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return nil, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetAllIssuedESDTs("")
	require.Equal(t, localErr, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetAllIssuedESDTsGetRestEndPointError(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			return 0, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetAllIssuedESDTs("")
	require.Equal(t, ErrSendingRequest, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetAllIssuedESDTs(t *testing.T) {
	t.Parallel()

	tokens := []string{"ESDT-5t6y7u", "NFT-9i8u7y-03"}
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			genericResp := &data.GenericAPIResponse{Data: tokens}
			genRespBytes, _ := json.Marshal(genericResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	genericResponse, err := nodeStatusProc.GetAllIssuedESDTs("")
	require.Nil(t, err)
	require.NotNil(t, genericResponse)

	slice, ok := genericResponse.Data.([]interface{})
	require.True(t, ok)

	for _, el := range slice {
		found := false
		for _, token := range tokens {
			if el.(string) == token {
				found = true
				break
			}
		}
		require.True(t, found)
	}
}

func TestNodeStatusProcessor_ApiPathIsCorrect(t *testing.T) {
	t.Parallel()

	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			require.Equal(t, path, "/network/esdt/semi-fungible-tokens")
			return 0, nil
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	_, err := nodeStatusProc.GetAllIssuedESDTs(data.SemiFungibleTokens)
	require.Nil(t, err)
}

func TestNodeStatusProcessor_GetDelegatedInfoGetObserversFailedShouldErr(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return nil, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetDelegatedInfo()
	require.Equal(t, localErr, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetDelegatedInfoGetRestEndPointError(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			return 0, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetDelegatedInfo()
	require.Equal(t, ErrSendingRequest, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetDelegatedInfo(t *testing.T) {
	t.Parallel()

	expectedResp := &data.GenericAPIResponse{Data: "delegated info"}
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			genRespBytes, _ := json.Marshal(expectedResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	actualResponse, err := nodeStatusProc.GetDelegatedInfo()
	require.Nil(t, err)
	require.Equal(t, expectedResp, actualResponse)
}

func TestNodeStatusProcessor_GetDirectStakedInfoGetObserversFailedShouldErr(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return nil, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetDirectStakedInfo()
	require.Equal(t, localErr, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetDirectStakedInfoGetRestEndPointError(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			return 0, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetDirectStakedInfo()
	require.Equal(t, ErrSendingRequest, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetDirectStakedInfo(t *testing.T) {
	t.Parallel()

	expectedResp := &data.GenericAPIResponse{Data: "direct staked info"}
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetObserversCalled: func(shardId uint32) (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			genRespBytes, _ := json.Marshal(expectedResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	actualResponse, err := nodeStatusProc.GetDirectStakedInfo()
	require.Nil(t, err)
	require.Equal(t, expectedResp, actualResponse)
}

func TestNodeStatusProcessor_GetEnableEpochsMetricsGetEndpointErr(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodesStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() ([]*data.NodeData, error) {
			return []*data.NodeData{
				{Address: "addr1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			return 0, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodesStatusProc.GetEnableEpochsMetrics()
	require.Equal(t, ErrSendingRequest, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetEnableEpochsMetricsShouldWork(t *testing.T) {
	t.Parallel()

	key := "smart_contract_deploy"
	expectedValue := float64(4)
	nodesStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() ([]*data.NodeData, error) {
			return []*data.NodeData{
				{Address: "addr1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			metricMap := map[string]interface{}{
				key: expectedValue,
			}
			genericResp := &data.GenericAPIResponse{Data: metricMap}
			genericRespBytes, _ := json.Marshal(genericResp)

			return 0, json.Unmarshal(genericRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	genericResponse, err := nodesStatusProc.GetEnableEpochsMetrics()
	require.Nil(t, err)
	require.NotNil(t, genericResponse)

	metricsMap, ok := genericResponse.Data.(map[string]interface{})
	require.True(t, ok)

	actualValue, ok := metricsMap[key]
	require.True(t, ok)
	require.Equal(t, expectedValue, actualValue)
}

func TestNodeStatusProcessor_GetEnableEpochsMetricsGetObserversShouldErr(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() ([]*data.NodeData, error) {
			return nil, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetEnableEpochsMetrics()
	require.Equal(t, localErr, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetRatingsConfigGetAllObserversShouldFail(t *testing.T) {
	t.Parallel()

	localErr := errors.New("local error")
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() ([]*data.NodeData, error) {
			return nil, localErr
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	status, err := nodeStatusProc.GetRatingsConfig()
	require.Equal(t, localErr, err)
	require.Nil(t, status)
}

func TestNodeStatusProcessor_GetRatingsConfig(t *testing.T) {
	t.Parallel()

	expectedResp := &data.GenericAPIResponse{Data: "ratings config"}
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			genRespBytes, _ := json.Marshal(expectedResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	actualResponse, err := nodeStatusProc.GetRatingsConfig()
	require.Nil(t, err)
	require.Equal(t, expectedResp, actualResponse)
}

func TestNodeStatusProcessor_GetGenesisNodesPubKeys(t *testing.T) {
	t.Parallel()

	expectedResp := &data.GenericAPIResponse{Data: "genesis nodes pub keys"}
	nodeStatusProc, _ := NewNodeStatusProcessor(&mock.ProcessorStub{
		GetAllObserversCalled: func() (observers []*data.NodeData, err error) {
			return []*data.NodeData{
				{Address: "address1", ShardId: 0},
			}, nil
		},
		CallGetRestEndPointCalled: func(address string, path string, value interface{}) (int, error) {
			genRespBytes, _ := json.Marshal(expectedResp)

			return 0, json.Unmarshal(genRespBytes, value)
		},
	},
		&mock.GenericApiResponseCacherMock{},
		time.Nanosecond,
	)

	actualResponse, err := nodeStatusProc.GetGenesisNodesPubKeys()
	require.Nil(t, err)
	require.Equal(t, expectedResp, actualResponse)
}
