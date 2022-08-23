package yolsdk

import (
	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/interactors"
)

func GetYolAddress(mnemonic string, walletIndex int64) (yollloAddr string, err error) {
	w := interactors.NewWallet()
	privateKey := w.GetPrivateKeyFromMnemonic(data.Mnemonic(mnemonic), 0, uint32(walletIndex))
	address, err := w.GetAddressFromPrivateKey(privateKey)
	if err != nil {

		return
	}

	return address.AddressAsBech32String(), nil
}

func GetPrivatKey64(mnemonic string, walletIndex int64) (privateKey64 []byte, err error) {
	w := interactors.NewWallet()
	privateKey := w.GetPrivateKeyFromMnemonic(data.Mnemonic(mnemonic), 0, uint32(walletIndex))
	address, err := w.GetAddressFromPrivateKey(privateKey)
	if err != nil {

		return
	}
	if len(privateKey) == 32 {
		privateKey = append(privateKey, address.AddressBytes()...)
	}

	return privateKey, nil
}
