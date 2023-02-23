package core

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"yolllo-manager/models"
	"yolllo-manager/pkg/yolsdk"
)

func (c *Core) CreateUserTransaction(req models.CreateUserTransactionReq) (resp models.CreateUserTransactionResp, err error) {
	senderWalletIndex, err := c.Repo.PG.GetWalletIndexByWalletAddress(req.SenderAddress)
	if err != nil {

		return
	}

	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + req.SenderAddress)
	if err != nil {

		return
	}
	defer resHTTP.Body.Close()
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var walletInfo models.GetAddressResp
	err = json.Unmarshal(body, &walletInfo)
	if err != nil {

		return
	}

	if walletInfo.Error != "" {
		err = errors.New(walletInfo.Error)

		return
	}

	// send trx
	var trxReq models.ProxyAPITransactionSendReq
	trxReq.Nonce = walletInfo.Data.Account.Nonce
	trxReq.Value = req.Value
	trxReq.Receiver = req.ReceiverAddress
	trxReq.Sender = walletInfo.Data.Account.Address
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 50000
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1
	signData := `{"nonce":` + strconv.FormatInt(trxReq.Nonce, 10) +
		`,"value":"` + trxReq.Value +
		`","receiver":"` + trxReq.Receiver +
		`","sender":"` + trxReq.Sender +
		`","gasPrice":` + strconv.FormatInt(trxReq.GasPrice, 10) +
		`,"gasLimit":` + strconv.FormatInt(trxReq.GasLimit, 10) +
		`,"chainID":"` + trxReq.ChainID +
		`","version":` + strconv.FormatInt(trxReq.Version, 10) + `}`

	userPrivateKey64, err := yolsdk.GetPrivatKey64(c.Config.Mnemonic, senderWalletIndex)
	if err != nil {

		return
	}
	sign := ed25519.Sign(ed25519.PrivateKey(userPrivateKey64), []byte(signData))
	trxReq.Signature = hex.EncodeToString(sign)

	jsonData, err := json.Marshal(trxReq)
	if err != nil {

		return
	}
	resHTTP, err = http.Post("http://"+c.Config.ProxyAddress+"/transaction/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err = ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var transactionInfo models.ProxyAPITransactionSendResp
	err = json.Unmarshal(body, &transactionInfo)
	if err != nil {

		return
	}

	if transactionInfo.Error != "" {
		err = errors.New(transactionInfo.Error)

		return
	}

	resp.TransactionHash = transactionInfo.Data.TxHash

	return
}

func (c *Core) GetTransaction(req models.GetTransactionReq) (resp models.GetTransactionResp, err error) {
	transactions, err := c.Repo.ES.GetTransactionByHash(req.TransactionHash)
	if err != nil {

		return
	}
	for _, transaction := range transactions.Hits.Hits {
		resp.Hash = transaction.ID
		resp.Nonce = transaction.Source.Nonce
		resp.Receiver = transaction.Source.Receiver
		resp.Sender = transaction.Source.Sender
		resp.ReceiverShard = transaction.Source.ReceiverShard
		resp.SenderShard = transaction.Source.SenderShard
		resp.Value = transaction.Source.Value
		resp.Timestamp = transaction.Source.Timestamp
		resp.Status = transaction.Source.Status
		break
	}

	return
}

func (c *Core) GetTransactionCost(req models.GetTransactionCostReq) (resp models.GetTransactionCostResp, err error) {
	var trxCostReq models.ProxyAPITransactionCostReq
	trxCostReq.Value = req.Value
	trxCostReq.Receiver = req.ReceiverAddress
	trxCostReq.Sender = req.SenderAddress
	trxCostReq.ChainID = "yolllo-network"
	trxCostReq.Version = 1

	jsonData, err := json.Marshal(trxCostReq)
	if err != nil {

		return
	}

	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/transaction/cost", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {

		return
	}
	if resp.Error != "" {
		err = errors.New(resp.Error)

		return
	}

	return
}

func (c *Core) GetTransactionFee(req models.GetTransactionFeeReq) (resp models.GetTransactionFeeResp, err error) {
	var trxCostReq models.ProxyAPITransactionCostReq
	trxCostReq.Value = req.Value
	trxCostReq.Receiver = req.ReceiverAddress
	trxCostReq.Sender = req.SenderAddress
	trxCostReq.ChainID = "yolllo-network"
	trxCostReq.Version = 1

	jsonData, err := json.Marshal(trxCostReq)
	if err != nil {

		return
	}

	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/transaction/cost", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var trxCostResp models.ProxyAPITransactionCostResp
	err = json.Unmarshal(body, &trxCostResp)
	if err != nil {

		return
	}
	if trxCostResp.Error != "" {
		err = errors.New(trxCostResp.Error)

		return
	}

	var gasPrice int64 = 1000000000
	resp.Value = big.NewInt(0).Mul(big.NewInt(trxCostResp.Data.TxGasUnits), big.NewInt(gasPrice))

	return
}

func (c *Core) CreateTransaction(req models.CreateTransactionReq) (resp models.CreateTransactionResp, err error) {
	var trxReq models.ProxyAPITransactionSendReq
	trxReq.Nonce = req.Nonce
	trxReq.Value = req.Value
	trxReq.Receiver = req.ReceiverAddress
	trxReq.Sender = req.SenderAddress
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 50000
	trxReq.Signature = req.Signature
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1

	jsonData, err := json.Marshal(trxReq)
	if err != nil {

		return
	}
	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/transaction/send", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var transactionInfo models.ProxyAPITransactionSendResp
	err = json.Unmarshal(body, &transactionInfo)
	if err != nil {

		return
	}

	if transactionInfo.Error != "" {
		err = errors.New(transactionInfo.Error)

		return
	}

	resp.TransactionHash = transactionInfo.Data.TxHash

	return
}

func (c *Core) GetLastTransactionList(req models.GetLastTransactionListReq) (resp models.GetLastTransactionListResp, err error) {
	transactions, err := c.Repo.ES.GetLastTransactions(req.PageSize)
	if err != nil {

		return
	}
	for _, transaction := range transactions.Hits.Hits {
		var transactionInfo models.TransactionListTransactionInfo
		transactionInfo.Hash = transaction.ID
		transactionInfo.Nonce = transaction.Source.Nonce
		transactionInfo.Receiver = transaction.Source.Receiver
		transactionInfo.Sender = transaction.Source.Sender
		transactionInfo.ReceiverShard = transaction.Source.ReceiverShard
		transactionInfo.SenderShard = transaction.Source.SenderShard
		transactionInfo.Value = transaction.Source.Value
		transactionInfo.Timestamp = transaction.Source.Timestamp
		transactionInfo.Status = transaction.Source.Status
		resp.TransactionList = append(resp.TransactionList, transactionInfo)
		if len(transaction.Sort) > 0 {
			resp.NextTimestampAfter = transaction.Sort[0]
			resp.NextSearchOrderAfter = transaction.Sort[1]
		}
	}

	return
}

func (c *Core) GetNextTransactionList(req models.GetNextTransactionListReq) (resp models.GetNextTransactionListResp, err error) {
	transactions, err := c.Repo.ES.GetNextTransactions(req.PageSize, req.TimestampAfter, req.SearchOrderAfter)
	if err != nil {

		return
	}
	for _, transaction := range transactions.Hits.Hits {
		var transactionInfo models.TransactionListTransactionInfo
		transactionInfo.Hash = transaction.ID
		transactionInfo.Nonce = transaction.Source.Nonce
		transactionInfo.Receiver = transaction.Source.Receiver
		transactionInfo.Sender = transaction.Source.Sender
		transactionInfo.ReceiverShard = transaction.Source.ReceiverShard
		transactionInfo.SenderShard = transaction.Source.SenderShard
		transactionInfo.Value = transaction.Source.Value
		transactionInfo.Timestamp = transaction.Source.Timestamp
		transactionInfo.Status = transaction.Source.Status
		resp.TransactionList = append(resp.TransactionList, transactionInfo)
		if len(transaction.Sort) > 0 {
			resp.NextTimestampAfter = transaction.Sort[0]
			resp.NextSearchOrderAfter = transaction.Sort[1]
		}
	}

	return
}

func (c *Core) GetLastTransactionListByAddr(req models.GetLastTransactionListByAddrReq) (resp models.GetLastTransactionListByAddrResp, err error) {
	transactions, err := c.Repo.ES.GetLastTransactionsByAddr(req.PageSize, req.WalletAddress)
	if err != nil {

		return
	}
	for _, transaction := range transactions.Hits.Hits {
		var transactionInfo models.TransactionListTransactionInfo
		transactionInfo.Hash = transaction.ID
		transactionInfo.Nonce = transaction.Source.Nonce
		transactionInfo.Receiver = transaction.Source.Receiver
		transactionInfo.Sender = transaction.Source.Sender
		transactionInfo.ReceiverShard = transaction.Source.ReceiverShard
		transactionInfo.SenderShard = transaction.Source.SenderShard
		transactionInfo.Value = transaction.Source.Value
		transactionInfo.Timestamp = transaction.Source.Timestamp
		transactionInfo.Status = transaction.Source.Status
		resp.TransactionList = append(resp.TransactionList, transactionInfo)
		if len(transaction.Sort) > 0 {
			resp.NextTimestampAfter = transaction.Sort[0]
			resp.NextSearchOrderAfter = transaction.Sort[1]
		}
	}

	return
}

func (c *Core) GetNextTransactionListByAddr(req models.GetNextTransactionListByAddrReq) (resp models.GetNextTransactionListByAddrResp, err error) {
	transactions, err := c.Repo.ES.GetNextTransactionsByAddr(req.PageSize, req.WalletAddress, req.TimestampAfter, req.SearchOrderAfter)
	if err != nil {

		return
	}
	for _, transaction := range transactions.Hits.Hits {
		var transactionInfo models.TransactionListTransactionInfo
		transactionInfo.Hash = transaction.ID
		transactionInfo.Nonce = transaction.Source.Nonce
		transactionInfo.Receiver = transaction.Source.Receiver
		transactionInfo.Sender = transaction.Source.Sender
		transactionInfo.ReceiverShard = transaction.Source.ReceiverShard
		transactionInfo.SenderShard = transaction.Source.SenderShard
		transactionInfo.Value = transaction.Source.Value
		transactionInfo.Timestamp = transaction.Source.Timestamp
		transactionInfo.Status = transaction.Source.Status
		resp.TransactionList = append(resp.TransactionList, transactionInfo)
		if len(transaction.Sort) > 0 {
			resp.NextTimestampAfter = transaction.Sort[0]
			resp.NextSearchOrderAfter = transaction.Sort[1]
		}
	}

	return
}

func (c *Core) GetRangeTransactionList(req models.GetRangeTransactionListReq) (resp models.GetRangeTransactionListResp, err error) {
	transactions, err := c.Repo.ES.GetRangeTransactions(req.PageSize, req.PageFrom, req.TimestampFrom, req.TimestampTo)
	if err != nil {

		return
	}
	for _, transaction := range transactions.Hits.Hits {
		var transactionInfo models.TransactionListTransactionInfo
		transactionInfo.Hash = transaction.ID
		transactionInfo.Nonce = transaction.Source.Nonce
		transactionInfo.Receiver = transaction.Source.Receiver
		transactionInfo.Sender = transaction.Source.Sender
		transactionInfo.ReceiverShard = transaction.Source.ReceiverShard
		transactionInfo.SenderShard = transaction.Source.SenderShard
		transactionInfo.Value = transaction.Source.Value
		transactionInfo.Timestamp = transaction.Source.Timestamp
		transactionInfo.Status = transaction.Source.Status
		resp.TransactionList = append(resp.TransactionList, transactionInfo)
	}
	resp.Total = transactions.Hits.Total.Value

	return
}
