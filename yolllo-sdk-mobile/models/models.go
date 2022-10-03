package models

type GetWalletReq struct {
	Mnemonic string `json:"mnemonic"`
	Index    uint32 `json:"index"`
}

type GetSignatureReq struct {
	Nonce      int64  `json:"nonce"`
	Value      string `json:"value"`
	Receiver   string `json:"receiver"`
	Sender     string `json:"sender"`
	PrivateKey string `json:"private_key"`
}
