//gomobile bind -target android -o android/yolllosdk.aar
package yollosdkmobile

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"yolllo-sdk-mobile/models"

	"github.com/ElrondNetwork/elrond-sdk-erdgo/data"
	"github.com/ElrondNetwork/elrond-sdk-erdgo/interactors"
)

/*
GetWallet()
req:
{
	"mnemonic":		string,
	"index":		integer,
}
resp:
{
	"wallet_address":	string,
	"private_key":		string,
	"error":			string
}
*/
func GetWallet(reqJSON string) string {
	var req models.GetWalletReq
	err := json.Unmarshal([]byte(reqJSON), &req)
	if err != nil {
		return GetWalletResp(
			"",
			"",
			err.Error(),
		)
	}

	w := interactors.NewWallet()
	privateKey := w.GetPrivateKeyFromMnemonic(data.Mnemonic(req.Mnemonic), 0, req.Index)
	address, err := w.GetAddressFromPrivateKey(privateKey)
	if err != nil {
		return GetWalletResp(
			"",
			"",
			err.Error(),
		)
	}
	if len(privateKey) == 32 {
		privateKey = append(privateKey, address.AddressBytes()...)
	}

	return GetWalletResp(
		address.AddressAsBech32String(),
		hex.EncodeToString(privateKey),
		"",
	)
}

func GetWalletResp(walletAddress, privateKey, err string) string {
	return fmt.Sprintf(`{"wallet_address":"%s","private_key":"%s","error":"%s"}`, walletAddress, privateKey, err)
}

/*
GetSignature()
req:
{
	"nonce":		integer,
	"value":		string,
	"receiver":		string,
	"sender":		string,
	"private_key":	string
}
resp:
{
	"signature":	string,
	"error":		string
}
*/
func GetSignature(reqJSON string) string {
	var req models.GetSignatureReq
	err := json.Unmarshal([]byte(reqJSON), &req)
	if err != nil {
		return GetSignatureResp(
			"",
			err.Error(),
		)
	}
	privKey, err := hex.DecodeString(string(req.PrivateKey))
	if err != nil {
		return GetSignatureResp(
			"",
			err.Error(),
		)
	}
	signData := `{"nonce":` + strconv.FormatInt(req.Nonce, 10) +
		`,"value":"` + req.Value +
		`","receiver":"` + req.Receiver +
		`","sender":"` + req.Sender +
		`","gasPrice":1000000000,"gasLimit":50000,"chainID":"yolllo-network","version":1}`
	signature := ed25519.Sign(ed25519.PrivateKey([]byte(privKey)), []byte(signData))

	return GetSignatureResp(
		hex.EncodeToString(signature),
		"",
	)
}

func GetSignatureResp(signature, err string) string {
	return fmt.Sprintf(`{"signature":"%s","error":"%s"}`, signature, err)
}
