package blockchain

import (
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go/sharding"
	erdgoCore "github.com/ElrondNetwork/elrond-sdk-erdgo/core"
)

type shardCoordinator struct {
	coordinator sharding.Coordinator
}

// NewShardCoordinator returns a shard coordinator instance that is able to execute sharding-related operations
func NewShardCoordinator(numOfShardsWithoutMeta uint32, currentShard uint32) (*shardCoordinator, error) {
	coord, err := sharding.NewMultiShardCoordinator(numOfShardsWithoutMeta, currentShard)
	if err != nil {
		return nil, err
	}

	return &shardCoordinator{
		coordinator: coord,
	}, nil
}

// ComputeShardId computes the shard ID of a provided address
func (sc *shardCoordinator) ComputeShardId(address erdgoCore.AddressHandler) (uint32, error) {
	if check.IfNil(address) {
		return 0, ErrNilAddress
	}
	if len(address.AddressBytes()) == 0 {
		return 0, ErrInvalidAddress
	}

	return sc.coordinator.ComputeId(address.AddressBytes()), nil
}

// IsInterfaceNil returns true if there is no value under the interface
func (sc *shardCoordinator) IsInterfaceNil() bool {
	return sc == nil
}
