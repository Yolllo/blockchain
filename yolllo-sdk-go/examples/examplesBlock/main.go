package main

import (
	"context"
	"encoding/json"
	"time"

	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/blockchain"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/core"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/examples"
)

var log = logger.GetOrCreate("elrond-sdk-erdgo/examples/examplesBlock")

func main() {
	args := blockchain.ArgsElrondProxy{
		ProxyURL:            examples.TestnetGateway,
		Client:              nil,
		SameScState:         false,
		ShouldBeSynced:      false,
		FinalityCheck:       false,
		CacheExpirationTime: time.Minute,
		EntityType:          core.Proxy,
	}
	ep, err := blockchain.NewElrondProxy(args)
	if err != nil {
		log.Error("error creating proxy", "error", err)
		return
	}

	// Get latest hyper block (metachain) nonce
	nonce, err := ep.GetLatestHyperBlockNonce(context.Background())
	if err != nil {
		log.Error("error retrieving latest block nonce", "error", err)
		return
	}
	log.Info("latest hyper block", "nonce", nonce)

	// Get block info
	block, errGet := ep.GetHyperBlockByNonce(context.Background(), nonce)
	if errGet != nil {
		log.Error("error retrieving hyper block", "error", err)
		return
	}
	data, errMarshal := json.MarshalIndent(block, "", "    ")
	if errMarshal != nil {
		log.Error("error serializing block", "error", errMarshal)
		return
	}
	log.Info("\n" + string(data))
}
