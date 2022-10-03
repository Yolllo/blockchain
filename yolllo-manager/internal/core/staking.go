package core

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"yolllo-manager/models"
	"yolllo-manager/pkg/helper"
	"yolllo-manager/pkg/yolsdk"

	"github.com/ElrondNetwork/elrond-go-core/core/mock"
	"github.com/ElrondNetwork/elrond-go-core/core/pubkeyConverter"
)

func (c *Core) DelegateUserStaking(req models.DelegateUserStakingReq) (resp models.DelegateUserStakingResp, err error) {
	userWalletIndex, err := c.Repo.PG.GetWalletIndexByWalletAddress(req.UserAddress)
	if err != nil {

		return
	}

	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + req.UserAddress)
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
	var trxReq models.ProxyAPITransactionSendWithDataReq
	trxReq.Nonce = walletInfo.Data.Account.Nonce
	trxReq.Value = req.Value
	trxReq.Receiver = c.Config.StakingAddress
	trxReq.Sender = walletInfo.Data.Account.Address
	trxReq.Data = "ZGVsZWdhdGU=" // "delegate"
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 55099500
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1
	signData := `{"nonce":` + strconv.FormatInt(trxReq.Nonce, 10) +
		`,"value":"` + trxReq.Value +
		`","receiver":"` + trxReq.Receiver +
		`","sender":"` + trxReq.Sender +
		`","gasPrice":` + strconv.FormatInt(trxReq.GasPrice, 10) +
		`,"gasLimit":` + strconv.FormatInt(trxReq.GasLimit, 10) +
		`,"data":"` + trxReq.Data +
		`","chainID":"` + trxReq.ChainID +
		`","version":` + strconv.FormatInt(trxReq.Version, 10) + `}`

	userPrivateKey64, err := yolsdk.GetPrivatKey64(c.Config.Mnemonic, userWalletIndex)
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

func (c *Core) GetUserStaking(req models.GetUserStakingReq) (resp models.GetUserStakingResp, err error) {
	// get hex addr
	bech32, err := pubkeyConverter.NewBech32PubkeyConverter(32, &mock.LoggerMock{})
	if err != nil {

		return
	}
	userAddrHexByte, err := bech32.Decode(req.UserAddress)
	if err != nil {

		return
	}
	userAddrHex := hex.EncodeToString(userAddrHexByte)

	// query for get value
	var queryReq models.ProxyAPIQueryClaimableRewardsReq
	queryReq.SCAddress = c.Config.StakingAddress
	queryReq.FuncName = "getUserActiveStake"
	queryReq.Args = append(queryReq.Args, userAddrHex)
	jsonData, err := json.Marshal(queryReq)
	if err != nil {

		return
	}
	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/vm-values/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var queryResp models.ProxyAPIQueryClaimableRewardsResp
	err = json.Unmarshal(body, &queryResp)
	if err != nil {

		return
	}

	if queryResp.Error != "" {
		err = errors.New(queryResp.Error)

		return
	}

	if len(queryResp.Data.Data.ReturnData) == 0 {
		resp.Value = "0"

		return
	}

	// decode value
	rewardHexByte, err := base64.StdEncoding.DecodeString(queryResp.Data.Data.ReturnData[0])
	if err != nil {

		return
	}
	rewardHex := hex.EncodeToString(rewardHexByte)
	rewardBigInt := new(big.Int)
	rewardBigInt.SetString(rewardHex, 16)

	resp.Value = rewardBigInt.String()

	return
}

func (c *Core) UndelegateUserStaking(req models.UndelegateUserStakingReq) (resp models.UndelegateUserStakingResp, err error) {
	userWalletIndex, err := c.Repo.PG.GetWalletIndexByWalletAddress(req.UserAddress)
	if err != nil {

		return
	}

	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + req.UserAddress)
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
	var trxReq models.ProxyAPITransactionSendWithDataReq
	trxReq.Nonce = walletInfo.Data.Account.Nonce
	trxReq.Value = "0"
	trxReq.Receiver = c.Config.StakingAddress
	trxReq.Sender = walletInfo.Data.Account.Address
	valueBigInt := new(big.Int)
	valueBigInt, ok := valueBigInt.SetString(req.Value, 10)
	if !ok {
		err = errors.New("SetString: error")
		return
	}
	dataStr := "unDelegate@" + helper.BigintToHex(valueBigInt)

	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 12000000
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1
	signData := `{"nonce":` + strconv.FormatInt(trxReq.Nonce, 10) +
		`,"value":"` + trxReq.Value +
		`","receiver":"` + trxReq.Receiver +
		`","sender":"` + trxReq.Sender +
		`","gasPrice":` + strconv.FormatInt(trxReq.GasPrice, 10) +
		`,"gasLimit":` + strconv.FormatInt(trxReq.GasLimit, 10) +
		`,"data":"` + trxReq.Data +
		`","chainID":"` + trxReq.ChainID +
		`","version":` + strconv.FormatInt(trxReq.Version, 10) + `}`

	userPrivateKey64, err := yolsdk.GetPrivatKey64(c.Config.Mnemonic, userWalletIndex)
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

func (c *Core) GetUserStakingUndelegated(req models.GetUserStakingUndelegatedReq) (resp models.GetUserStakingUndelegatedResp, err error) {
	// get hex addr
	bech32, err := pubkeyConverter.NewBech32PubkeyConverter(32, &mock.LoggerMock{})
	if err != nil {

		return
	}
	userAddrHexByte, err := bech32.Decode(req.UserAddress)
	if err != nil {

		return
	}
	userAddrHex := hex.EncodeToString(userAddrHexByte)

	// query for get value
	var queryReq models.ProxyAPIQueryClaimableRewardsReq
	queryReq.SCAddress = c.Config.StakingAddress
	queryReq.FuncName = "getUserUnDelegatedList"
	queryReq.Args = append(queryReq.Args, userAddrHex)
	jsonData, err := json.Marshal(queryReq)
	if err != nil {

		return
	}
	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/vm-values/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var queryResp models.ProxyAPIQueryGetUndelegatedListResp
	err = json.Unmarshal(body, &queryResp)
	if err != nil {

		return
	}

	if queryResp.Error != "" {
		err = errors.New(queryResp.Error)

		return
	}

	// decode value
	for _, valueBase64 := range queryResp.Data.Data.ReturnData {
		rewardHexByte, err := base64.StdEncoding.DecodeString(valueBase64)
		if err != nil {

			return resp, err
		}
		rewardHex := hex.EncodeToString(rewardHexByte)
		rewardBigInt := new(big.Int)
		rewardBigInt.SetString(rewardHex, 16)

		resp.Values = append(resp.Values, rewardBigInt.String())
	}

	return
}

func (c *Core) ClaimUserStakingUndelegated(req models.ClaimUserStakingUndelegatedReq) (resp models.ClaimUserStakingUndelegatedResp, err error) {
	userWalletIndex, err := c.Repo.PG.GetWalletIndexByWalletAddress(req.UserAddress)
	if err != nil {

		return
	}

	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + req.UserAddress)
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
	var trxReq models.ProxyAPITransactionSendWithDataReq
	trxReq.Nonce = walletInfo.Data.Account.Nonce
	trxReq.Value = "0"
	trxReq.Receiver = c.Config.StakingAddress
	trxReq.Sender = walletInfo.Data.Account.Address
	dataStr := "withdraw"
	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 12000000
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1
	signData := `{"nonce":` + strconv.FormatInt(trxReq.Nonce, 10) +
		`,"value":"` + trxReq.Value +
		`","receiver":"` + trxReq.Receiver +
		`","sender":"` + trxReq.Sender +
		`","gasPrice":` + strconv.FormatInt(trxReq.GasPrice, 10) +
		`,"gasLimit":` + strconv.FormatInt(trxReq.GasLimit, 10) +
		`,"data":"` + trxReq.Data +
		`","chainID":"` + trxReq.ChainID +
		`","version":` + strconv.FormatInt(trxReq.Version, 10) + `}`

	userPrivateKey64, err := yolsdk.GetPrivatKey64(c.Config.Mnemonic, userWalletIndex)
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

func (c *Core) GetUserStakingReward(req models.GetUserStakingRewardReq) (resp models.GetUserStakingRewardResp, err error) {
	// get hex addr
	bech32, err := pubkeyConverter.NewBech32PubkeyConverter(32, &mock.LoggerMock{})
	if err != nil {

		return
	}
	userAddrHexByte, err := bech32.Decode(req.UserAddress)
	if err != nil {

		return
	}
	userAddrHex := hex.EncodeToString(userAddrHexByte)

	// query for get value
	var queryReq models.ProxyAPIQueryClaimableRewardsReq
	queryReq.SCAddress = c.Config.StakingAddress
	queryReq.FuncName = "getClaimableRewards"
	queryReq.Args = append(queryReq.Args, userAddrHex)
	jsonData, err := json.Marshal(queryReq)
	if err != nil {

		return
	}
	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/vm-values/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var queryResp models.ProxyAPIQueryClaimableRewardsResp
	err = json.Unmarshal(body, &queryResp)
	if err != nil {

		return
	}

	if queryResp.Error != "" {
		err = errors.New(queryResp.Error)

		return
	}

	if len(queryResp.Data.Data.ReturnData) == 0 {
		resp.RewardValue = "0"

		return
	}
	fmt.Println(queryResp)
	// decode value
	rewardHexByte, err := base64.StdEncoding.DecodeString(queryResp.Data.Data.ReturnData[0])
	if err != nil {

		return
	}
	rewardHex := hex.EncodeToString(rewardHexByte)
	rewardBigInt := new(big.Int)
	rewardBigInt.SetString(rewardHex, 16)

	resp.RewardValue = rewardBigInt.String()

	return
}

func (c *Core) ClaimUserStakingReward(req models.ClaimUserStakingRewardReq) (resp models.ClaimUserStakingRewardResp, err error) {
	userWalletIndex, err := c.Repo.PG.GetWalletIndexByWalletAddress(req.UserAddress)
	if err != nil {

		return
	}

	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + req.UserAddress)
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
	var trxReq models.ProxyAPITransactionSendWithDataReq
	trxReq.Nonce = walletInfo.Data.Account.Nonce
	trxReq.Value = "0"
	trxReq.Receiver = c.Config.StakingAddress
	trxReq.Sender = walletInfo.Data.Account.Address
	trxReq.Data = "Y2xhaW1SZXdhcmRz" // "claimRewards"
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 6000000
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1
	signData := `{"nonce":` + strconv.FormatInt(trxReq.Nonce, 10) +
		`,"value":"` + trxReq.Value +
		`","receiver":"` + trxReq.Receiver +
		`","sender":"` + trxReq.Sender +
		`","gasPrice":` + strconv.FormatInt(trxReq.GasPrice, 10) +
		`,"gasLimit":` + strconv.FormatInt(trxReq.GasLimit, 10) +
		`,"data":"` + trxReq.Data +
		`","chainID":"` + trxReq.ChainID +
		`","version":` + strconv.FormatInt(trxReq.Version, 10) + `}`

	userPrivateKey64, err := yolsdk.GetPrivatKey64(c.Config.Mnemonic, userWalletIndex)
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

func (c *Core) GetUserStakingTotalStake() (resp models.GetUserStakingTotalStakeResp, err error) {
	// query for get value
	var queryReq models.ProxyAPIQueryGetTotalActiveStakeReq
	queryReq.SCAddress = c.Config.StakingAddress
	queryReq.FuncName = "getTotalActiveStake"
	jsonData, err := json.Marshal(queryReq)
	if err != nil {

		return
	}
	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/vm-values/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var queryResp models.ProxyAPIQueryGetTotalActiveStakeResp
	err = json.Unmarshal(body, &queryResp)
	if err != nil {

		return
	}

	if queryResp.Error != "" {
		err = errors.New(queryResp.Error)

		return
	}

	if len(queryResp.Data.Data.ReturnData) == 0 {
		resp.TotalStakeValue = "0"

		return
	}

	// decode value
	rewardHexByte, err := base64.StdEncoding.DecodeString(queryResp.Data.Data.ReturnData[0])
	if err != nil {

		return
	}
	rewardHex := hex.EncodeToString(rewardHexByte)
	rewardBigInt := new(big.Int)
	rewardBigInt.SetString(rewardHex, 16)

	resp.TotalStakeValue = rewardBigInt.String()

	return
}

func (c *Core) GetUserStakingFee() (resp models.GetUserStakingFeeResp, err error) {
	// query for get value
	var queryReq models.ProxyAPIQueryGetContractConfigReq
	queryReq.SCAddress = c.Config.StakingAddress
	queryReq.FuncName = "getContractConfig"
	jsonData, err := json.Marshal(queryReq)
	if err != nil {

		return
	}
	resHTTP, err := http.Post("http://"+c.Config.ProxyAddress+"/vm-values/query", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {

		return
	}
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var queryResp models.ProxyAPIQueryGetContractConfigResp
	err = json.Unmarshal(body, &queryResp)
	if err != nil {

		return
	}

	if queryResp.Error != "" {
		err = errors.New(queryResp.Error)

		return
	}

	if len(queryResp.Data.Data.ReturnData) != 10 {
		err = errors.New("unknown error")
		return
	}

	feeBase64 := queryResp.Data.Data.ReturnData[1]
	if feeBase64 == "" {
		resp.Value = float64(0)
		return
	}

	fmt.Println(feeBase64)

	feeHexByte, err := base64.StdEncoding.DecodeString(feeBase64)
	if err != nil {
		return
	}

	feeHex := hex.EncodeToString(feeHexByte)
	feeInt, err := strconv.ParseInt(feeHex, 16, 64)
	if err != nil {
		return
	}

	resp.Value = float64(feeInt) / 100

	return
}

func (c *Core) SetUserStakingFee(req models.SetUserStakingFeeReq) (resp models.SetUserStakingFeeResp, err error) {
	userWalletIndex, err := c.Repo.PG.GetWalletIndexByWalletAddress(c.Config.StakingOwner)
	if err != nil {

		return
	}

	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + c.Config.StakingOwner)
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

	// fee calc
	feeUnfloated := int64(req.Value * 100)
	feeHexStr := helper.Int64ToHex(feeUnfloated)

	// send trx
	var trxReq models.ProxyAPITransactionSendWithDataReq
	trxReq.Nonce = walletInfo.Data.Account.Nonce
	trxReq.Value = "0"
	trxReq.Receiver = c.Config.StakingAddress
	trxReq.Sender = walletInfo.Data.Account.Address
	dataStr := "changeServiceFee@" + feeHexStr
	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000
	trxReq.GasLimit = 55099500
	trxReq.ChainID = "yolllo-network"
	trxReq.Version = 1
	signData := `{"nonce":` + strconv.FormatInt(trxReq.Nonce, 10) +
		`,"value":"` + trxReq.Value +
		`","receiver":"` + trxReq.Receiver +
		`","sender":"` + trxReq.Sender +
		`","gasPrice":` + strconv.FormatInt(trxReq.GasPrice, 10) +
		`,"gasLimit":` + strconv.FormatInt(trxReq.GasLimit, 10) +
		`,"data":"` + trxReq.Data +
		`","chainID":"` + trxReq.ChainID +
		`","version":` + strconv.FormatInt(trxReq.Version, 10) + `}`

	userPrivateKey64, err := yolsdk.GetPrivatKey64(c.Config.Mnemonic, userWalletIndex)
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

func (c *Core) GetStakingCurrentMonthlyReward() (resp models.GetStakingCurrentMonthlyRewardResp, err error) {
	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/network/status/0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resHTTP.Body.Close()
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	var shardInfo models.ProxyAPINetworkStatusShardResp
	err = json.Unmarshal(body, &shardInfo)
	if err != nil {
		fmt.Println(shardInfo)
		err = errors.New("unknown error")

		return
	}

	if shardInfo.Error != "" {
		if shardInfo.Error == "the specified shard ID does not exist in proxy's configuration" {
			err = errors.New(shardInfo.Error)
		} else {
			fmt.Println(shardInfo)
			err = errors.New("unknown error")
		}

		return
	}

	monthIndex := getMonthIndex(483540)
	denomination := big.NewInt(1000000000000000000)
	var rewardPerMonthT int64
	switch monthIndex {
	case 1:
		rewardPerMonthT = 2000000
	case 2:
		rewardPerMonthT = 1950000
	case 3:
		rewardPerMonthT = 1901250
	case 4:
		rewardPerMonthT = 1853719
	case 5:
		rewardPerMonthT = 1807376
	case 6:
		rewardPerMonthT = 1762192
	case 7:
		rewardPerMonthT = 1718137
	case 8:
		rewardPerMonthT = 1675184
	case 9:
		rewardPerMonthT = 1633304
	case 10:
		rewardPerMonthT = 1592472
	case 11:
		rewardPerMonthT = 1552660
	case 12:
		rewardPerMonthT = 1513843
	case 13:
		rewardPerMonthT = 1475997
	case 14:
		rewardPerMonthT = 1439097
	case 15:
		rewardPerMonthT = 1403120
	case 16:
		rewardPerMonthT = 1368042
	case 17:
		rewardPerMonthT = 1333841
	case 18:
		rewardPerMonthT = 1300495
	case 19:
		rewardPerMonthT = 1267982
	case 20:
		rewardPerMonthT = 1236283
	case 21:
		rewardPerMonthT = 1205376
	case 22:
		rewardPerMonthT = 1175241
	case 23:
		rewardPerMonthT = 1145860
	case 24:
		rewardPerMonthT = 1117214
	case 25:
		rewardPerMonthT = 1089284
	case 26:
		rewardPerMonthT = 1062052
	case 27:
		rewardPerMonthT = 1035500
	case 28:
		rewardPerMonthT = 1009613
	case 29:
		rewardPerMonthT = 984372
	case 30:
		rewardPerMonthT = 959763
	case 31:
		rewardPerMonthT = 935769
	case 32:
		rewardPerMonthT = 912375
	case 33:
		rewardPerMonthT = 889566
	case 34:
		rewardPerMonthT = 867326
	case 35:
		rewardPerMonthT = 845643
	case 36:
		rewardPerMonthT = 824502
	case 37:
		rewardPerMonthT = 803890
	case 38:
		rewardPerMonthT = 783792
	case 39:
		rewardPerMonthT = 764198
	case 40:
		rewardPerMonthT = 745093
	case 41:
		rewardPerMonthT = 726465
	case 42:
		rewardPerMonthT = 708304
	case 43:
		rewardPerMonthT = 690596
	case 44:
		rewardPerMonthT = 673331
	case 45:
		rewardPerMonthT = 656498
	case 46:
		rewardPerMonthT = 640086
	case 47:
		rewardPerMonthT = 624083
	case 48:
		rewardPerMonthT = 608481
	case 49:
		rewardPerMonthT = 593269
	case 50:
		rewardPerMonthT = 578438
	case 51:
		rewardPerMonthT = 563977
	case 52:
		rewardPerMonthT = 549877
	case 53:
		rewardPerMonthT = 536130
	case 54:
		rewardPerMonthT = 522727
	case 55:
		rewardPerMonthT = 509659
	case 56:
		rewardPerMonthT = 496918
	case 57:
		rewardPerMonthT = 484495
	case 58:
		rewardPerMonthT = 472382
	case 59:
		rewardPerMonthT = 460573
	case 60:
		rewardPerMonthT = 449058

	}
	resp.Value = big.NewInt(0).Mul(big.NewInt(rewardPerMonthT), denomination)

	return
}

func getMonthIndex(currentRound int64) uint32 {
	const numberOfDaysInMounth = 28
	const numberOfSecondsInDay = 86400
	const roundTime = 5
	roundsPerDay := numberOfSecondsInDay / int64(roundTime)
	roundsPerMonth := numberOfDaysInMounth * roundsPerDay
	monthIndex := uint32(currentRound/roundsPerMonth) + 1
	if monthIndex > 60 {
		monthIndex = 60
	}
	return monthIndex
}

func (c *Core) IsValidAddress(req models.IsValidAddressReq) (resp models.IsValidAddressResp, err error) {
	bech32, err := pubkeyConverter.NewBech32PubkeyConverter(32, &mock.LoggerMock{})
	if err != nil {

		return
	}
	_, err = bech32.Decode(req.WalletAddress)
	if err != nil {
		resp.IsValid = false
		return resp, nil
	}

	resp.IsValid = true

	return resp, nil
}
