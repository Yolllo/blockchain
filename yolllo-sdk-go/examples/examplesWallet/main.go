package main

import (
	"fmt"

	"github.com/ElrondNetwork/elrond-go-crypto/signing"
	"github.com/ElrondNetwork/elrond-go-crypto/signing/mcl"
	logger "github.com/ElrondNetwork/elrond-go-logger"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/interactors"
)

var log = logger.GetOrCreate("elrond-sdk-erdgo/examples/examplesWallet")

func main() {
	w := interactors.NewWallet()
	mnemonic, err := w.GenerateMnemonic()
	if err != nil {
		log.Error("error generating mnemonic", "error", err)
		return
	}
	log.Info("generated mnemonics", "mnemonics", string(mnemonic))

	// generating the private key from the mnemonic using index 0
	index0 := uint32(0)
	privateKey0 := w.GetPrivateKeyFromMnemonic(mnemonic, 0, index0)
	address0, err := w.GetAddressFromPrivateKey(privateKey0)
	if err != nil {
		log.Error("error getting address from private key", "error", err)
		return
	}

	log.Info("generated private/public key",
		"private key", privateKey0,
		"index", index0,
		"address as hex", address0.AddressBytes(),
		"address as bech32", address0.AddressAsBech32String(),
	)

	// generating the private key from the same mnemonic using index 1
	index1 := uint32(1)
	privateKey1 := w.GetPrivateKeyFromMnemonic(mnemonic, 0, index1)
	address1, err := w.GetAddressFromPrivateKey(privateKey1)
	if err != nil {
		log.Error("error getting address from private key", "error", err)
		return
	}

	log.Info("generated private/public key",
		"private key", privateKey1,
		"index", index1,
		"address as hex", address1.AddressBytes(),
		"address as bech32", address1.AddressAsBech32String(),
	)

	validatorSuite := mcl.NewSuiteBLS12()
	validatorKeyGenerator := signing.NewKeyGenerator(validatorSuite)
	_, pk := validatorKeyGenerator.GeneratePair()
	pkByte, _ := pk.ToByteArray()

	fmt.Println(len(pkByte))

	/*
		var suite = ed25519.NewEd25519()
		var keyGenerator = signing.NewKeyGenerator(suite)
		privateKey, err := keyGenerator.PrivateKeyFromByteArray([]byte(privateKey0))
		publicKey := privateKey.GeneratePublic()
		publicKeyByte, _ := publicKey.ToByteArray()
		fmt.Println(len(publicKeyByte))
	*/
}
