package faucet

import (
	"encoding/hex"
	"strings"

	"github.com/ElrondNetwork/elrond-go-core/core"
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	crypto "github.com/ElrondNetwork/elrond-go-crypto"
	"github.com/ElrondNetwork/elrond-go-crypto/signing"
	"github.com/ElrondNetwork/elrond-go-crypto/signing/ed25519"
	"github.com/ElrondNetwork/elrond-go/sharding"
)

func getSuite() crypto.Suite {
	return ed25519.NewEd25519()
}

// PrivateKeysLoader will handle fetching keys pairs from the pem file
type PrivateKeysLoader struct {
	keyGen          crypto.KeyGenerator
	pemFileLocation string
	shardCoord      sharding.Coordinator
	pubKeyConverter core.PubkeyConverter
}

// NewPrivateKeysLoader will return a new instance of PrivateKeysLoader
func NewPrivateKeysLoader(
	shardCoord sharding.Coordinator,
	pemFileLocation string,
	pubKeyConverter core.PubkeyConverter,
) (*PrivateKeysLoader, error) {
	if check.IfNil(shardCoord) {
		return nil, ErrNilShardCoordinator
	}
	if !core.FileExists(pemFileLocation) {
		return nil, ErrFaucetPemFileDoesNotExist
	}
	if check.IfNil(pubKeyConverter) {
		return nil, ErrNilPubKeyConverter
	}

	keyGen := signing.NewKeyGenerator(getSuite())
	return &PrivateKeysLoader{
		keyGen:          keyGen,
		shardCoord:      shardCoord,
		pemFileLocation: pemFileLocation,
		pubKeyConverter: pubKeyConverter,
	}, nil
}

// PrivateKeysByShard will return a map containing private keys by shard ID
func (pkl *PrivateKeysLoader) PrivateKeysByShard() (map[uint32][]crypto.PrivateKey, error) {
	privKeysMapByShard := make(map[uint32][]crypto.PrivateKey)
	privKeysBytes, err := pkl.loadPrivKeysBytesFromPemFile()
	if err != nil {
		return nil, err
	}

	for _, privKeyBytes := range privKeysBytes {
		pkBytes, errD := hex.DecodeString(string(privKeyBytes))
		if errD != nil {
			return nil, errD
		}

		privKey, errP := pkl.keyGen.PrivateKeyFromByteArray(pkBytes)
		if errP != nil {
			return nil, errP
		}

		pubKeyOfPrivKey, errPk := pkl.pubKeyFromPrivKey(privKey)
		if errPk != nil {
			return nil, errPk
		}

		shardID := pkl.shardCoord.ComputeId(pubKeyOfPrivKey)

		privKeysMapByShard[shardID] = append(privKeysMapByShard[shardID], privKey)
	}

	return privKeysMapByShard, nil
}

func (pkl *PrivateKeysLoader) loadPrivKeysBytesFromPemFile() ([][]byte, error) {
	var privateKeysSlice [][]byte
	index := 0
	for {
		sk, _, err := core.LoadSkPkFromPemFile(pkl.pemFileLocation, index)
		if err != nil {
			if strings.Contains(err.Error(), "pem file is invalid") {
				return nil, err
			}

			if strings.Contains(err.Error(), "invalid private key index") {
				if len(privateKeysSlice) == 0 {
					return nil, err
				}

				return privateKeysSlice, nil
			}
		}

		privateKeysSlice = append(privateKeysSlice, sk)
		index++
	}
}

func (pkl *PrivateKeysLoader) pubKeyFromPrivKey(sk crypto.PrivateKey) ([]byte, error) {
	pk := sk.GeneratePublic()
	return pk.ToByteArray()
}
