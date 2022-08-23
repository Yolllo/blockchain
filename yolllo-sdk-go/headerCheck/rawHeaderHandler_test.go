package headerCheck_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
	"github.com/ElrondNetwork/elrond-go/state"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/headerCheck"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/testsCommon"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRawHeaderHandler(t *testing.T) {
	t.Parallel()

	t.Run("nil marshaller", func(t *testing.T) {
		t.Parallel()

		rh, err := headerCheck.NewRawHeaderHandler(&testsCommon.ProxyStub{}, nil)

		assert.True(t, check.IfNil(rh))
		assert.True(t, errors.Is(err, headerCheck.ErrNilMarshaller))
	})
	t.Run("nil proxy", func(t *testing.T) {
		t.Parallel()

		rh, err := headerCheck.NewRawHeaderHandler(nil, &testsCommon.MarshalizerStub{})

		assert.True(t, check.IfNil(rh))
		assert.True(t, errors.Is(err, headerCheck.ErrNilProxy))
	})
	t.Run("should work", func(t *testing.T) {
		t.Parallel()

		rh, err := headerCheck.NewRawHeaderHandler(&testsCommon.ProxyStub{}, &testsCommon.MarshalizerStub{})

		assert.False(t, check.IfNil(rh))
		assert.Nil(t, err)
	})
}

func TestGetMetaBlockByHash_ShouldFail(t *testing.T) {
	t.Parallel()

	t.Run("proxy error", func(t *testing.T) {
		expectedErr := errors.New("proxy err")
		proxy := &testsCommon.ProxyStub{
			GetRawBlockByHashCalled: func(shardId uint32, hash string) ([]byte, error) {
				return nil, expectedErr
			},
		}

		rh, err := headerCheck.NewRawHeaderHandler(proxy, &testsCommon.MarshalizerMock{})
		require.Nil(t, err)

		_, err = rh.GetMetaBlockByHash(context.Background(), "dummy")
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
	t.Run("marshaller error", func(t *testing.T) {
		expectedErr := errors.New("unmarshall err")
		marshaller := &testsCommon.MarshalizerStub{
			UnmarshalCalled: func(_ interface{}, _ []byte) error {
				return expectedErr
			},
		}

		rh, err := headerCheck.NewRawHeaderHandler(&testsCommon.ProxyStub{}, marshaller)
		require.Nil(t, err)

		_, err = rh.GetMetaBlockByHash(context.Background(), "dummy")
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestGetMetaBlockByHash_ShouldWork(t *testing.T) {
	t.Parallel()

	header := &block.MetaBlock{
		Nonce: 1,
		Epoch: 1,
	}
	headerBytes, _ := json.Marshal(header)

	proxy := &testsCommon.ProxyStub{
		GetRawBlockByHashCalled: func(shardId uint32, hash string) ([]byte, error) {
			return headerBytes, nil
		},
	}

	rh, err := headerCheck.NewRawHeaderHandler(proxy, &testsCommon.MarshalizerMock{})
	require.Nil(t, err)

	metaBlock, err := rh.GetMetaBlockByHash(context.Background(), "dummy")
	require.Nil(t, err)

	assert.Equal(t, metaBlock, header)
}

func TestGetShardBlockByHash_ShouldFail(t *testing.T) {
	t.Parallel()

	t.Run("proxy error", func(t *testing.T) {
		expectedErr := errors.New("proxy err")
		proxy := &testsCommon.ProxyStub{
			GetRawBlockByHashCalled: func(shardId uint32, hash string) ([]byte, error) {
				return nil, expectedErr
			},
		}

		rh, err := headerCheck.NewRawHeaderHandler(proxy, &testsCommon.MarshalizerMock{})
		require.Nil(t, err)

		_, err = rh.GetShardBlockByHash(context.Background(), 1, "dummy")
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
	t.Run("marshaller error", func(t *testing.T) {
		expectedErr := errors.New("unmarshall err")
		marshaller := &testsCommon.MarshalizerStub{
			UnmarshalCalled: func(_ interface{}, _ []byte) error {
				return expectedErr
			},
		}

		rh, err := headerCheck.NewRawHeaderHandler(&testsCommon.ProxyStub{}, marshaller)
		require.Nil(t, err)

		_, err = rh.GetShardBlockByHash(context.Background(), 1, "dummy")
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestGetShardBlockByHash_ShouldWork(t *testing.T) {
	t.Parallel()

	header := &block.Header{
		Nonce: 1,
		Epoch: 1,
	}
	headerBytes, _ := json.Marshal(header)

	proxy := &testsCommon.ProxyStub{
		GetRawBlockByHashCalled: func(shardId uint32, hash string) ([]byte, error) {
			return headerBytes, nil
		},
	}

	rh, err := headerCheck.NewRawHeaderHandler(proxy, &testsCommon.MarshalizerMock{})
	require.Nil(t, err)

	metaBlock, err := rh.GetShardBlockByHash(context.Background(), 1, "dummy")
	require.Nil(t, err)

	assert.Equal(t, metaBlock, header)
}

func TestGetValidatorsInfoPerEpoch_ShouldWork(t *testing.T) {
	t.Parallel()

	miniBlockHeaders := []block.MiniBlockHeader{
		{
			Hash:            []byte("hash1"),
			SenderShardID:   0,
			ReceiverShardID: 1,
			Type:            block.PeerBlock,
		},
	}

	expectedRandomness := []byte("prev rand seed")

	metaBlock := &block.MetaBlock{
		Nonce:            1,
		Epoch:            1,
		PrevRandSeed:     expectedRandomness,
		MiniBlockHeaders: miniBlockHeaders,
	}
	metaBlockBytes, _ := json.Marshal(metaBlock)

	vid := &state.ShardValidatorInfo{
		PublicKey: []byte("public key 1"),
		ShardId:   0,
	}
	vidBytes, _ := json.Marshal(vid)

	expectedValidatorsInfo := []*state.ShardValidatorInfo{vid}

	miniBlock := &block.MiniBlock{
		TxHashes:        [][]byte{vidBytes},
		ReceiverShardID: 0,
		SenderShardID:   0,
		Type:            block.PeerBlock,
	}
	miniBlockBytes, _ := json.Marshal(miniBlock)

	proxy := &testsCommon.ProxyStub{
		GetRawMiniBlockByHashCalled: func(shardId uint32, hash string, epoch uint32) ([]byte, error) {
			return miniBlockBytes, nil
		},
		GetRawStartOfEpochMetaBlockCalled: func(epoch uint32) ([]byte, error) {
			return metaBlockBytes, nil
		},
	}

	rh, err := headerCheck.NewRawHeaderHandler(proxy, &testsCommon.MarshalizerMock{})
	require.Nil(t, err)

	validatorInfo, randomness, err := rh.GetValidatorsInfoPerEpoch(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, expectedRandomness, randomness)
	assert.Equal(t, expectedValidatorsInfo, validatorInfo)
}
