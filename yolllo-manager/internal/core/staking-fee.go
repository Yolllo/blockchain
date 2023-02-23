package core

import (
	"encoding/base64"
	"errors"
	"math/big"
	"yolllo-manager/models"
	"yolllo-manager/pkg/helper"
)

const (
	GAS_FOR_MOVEMENT           = 50000
	GAS_PER_DATABYTE           = 1500
	GAS_FOR_EXECUTION_DELEGATE = 6000000
	GAS_FOR_EXECUTION_REWARD   = 1000000
)

func (c *Core) DelegateFeeUserStaking(req models.DelegateUserStakingReq) (resp models.DelegateFeeUserStakingResp, err error) {
	var trxReq models.ProxyAPITransactionSendWithDataReq
	dataStr := "delegate"
	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000

	// calculate fee
	totalFee := big.NewInt(0)

	gasForMovement := int64(GAS_FOR_MOVEMENT)
	feeForMovement := big.NewInt(0).Mul(big.NewInt(gasForMovement), big.NewInt(trxReq.GasPrice)) // gasForMovement * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForMovement)

	gasForData := int64(GAS_PER_DATABYTE * len(dataStr))
	feeForData := big.NewInt(0).Mul(big.NewInt(gasForData), big.NewInt(trxReq.GasPrice)) // gasForData * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForData)

	feeForExecutionWithoutModifier := big.NewInt(0).Mul(big.NewInt(GAS_FOR_EXECUTION_DELEGATE), big.NewInt(trxReq.GasPrice)) // gasForExecution * trxReq.GasPrice
	feeForExecution := big.NewInt(0).Div(feeForExecutionWithoutModifier, big.NewInt(100))
	totalFee = big.NewInt(0).Add(totalFee, feeForExecution)

	resp.Value = totalFee.String()

	return
}

func (c *Core) UndelegateFeeUserStaking(req models.UndelegateUserStakingReq) (resp models.UndelegateFeeUserStakingResp, err error) {
	var trxReq models.ProxyAPITransactionSendWithDataReq
	valueBigInt := new(big.Int)
	valueBigInt, ok := valueBigInt.SetString(req.Value, 10)
	if !ok {
		err = errors.New("SetString: error")
		return
	}
	dataStr := "unDelegate@" + helper.BigintToHex(valueBigInt)
	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000

	// calculate fee
	totalFee := big.NewInt(0)

	gasForMovement := int64(GAS_FOR_MOVEMENT)
	feeForMovement := big.NewInt(0).Mul(big.NewInt(gasForMovement), big.NewInt(trxReq.GasPrice)) // gasForMovement * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForMovement)

	gasForData := int64(GAS_PER_DATABYTE * len(dataStr))
	feeForData := big.NewInt(0).Mul(big.NewInt(gasForData), big.NewInt(trxReq.GasPrice)) // gasForData * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForData)

	feeForExecutionWithoutModifier := big.NewInt(0).Mul(big.NewInt(GAS_FOR_EXECUTION_DELEGATE), big.NewInt(trxReq.GasPrice)) // gasForExecution * trxReq.GasPrice
	feeForExecution := big.NewInt(0).Div(feeForExecutionWithoutModifier, big.NewInt(100))
	totalFee = big.NewInt(0).Add(totalFee, feeForExecution)

	resp.Value = totalFee.String()

	return
}

func (c *Core) ClaimFeeUserStakingUndelegated(req models.ClaimUserStakingUndelegatedReq) (resp models.ClaimFeeUserStakingUndelegatedResp, err error) {
	var trxReq models.ProxyAPITransactionSendWithDataReq
	dataStr := "withdraw"
	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000

	// calculate fee
	totalFee := big.NewInt(0)

	gasForMovement := int64(GAS_FOR_MOVEMENT)
	feeForMovement := big.NewInt(0).Mul(big.NewInt(gasForMovement), big.NewInt(trxReq.GasPrice)) // gasForMovement * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForMovement)

	gasForData := int64(GAS_PER_DATABYTE * len(dataStr))
	feeForData := big.NewInt(0).Mul(big.NewInt(gasForData), big.NewInt(trxReq.GasPrice)) // gasForData * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForData)

	feeForExecutionWithoutModifier := big.NewInt(0).Mul(big.NewInt(GAS_FOR_EXECUTION_DELEGATE), big.NewInt(trxReq.GasPrice)) // gasForExecution * trxReq.GasPrice
	feeForExecution := big.NewInt(0).Div(feeForExecutionWithoutModifier, big.NewInt(100))
	totalFee = big.NewInt(0).Add(totalFee, feeForExecution)

	resp.Value = totalFee.String()

	return
}

func (c *Core) ClaimFeeUserStakingReward(req models.ClaimUserStakingRewardReq) (resp models.ClaimFeeUserStakingRewardResp, err error) {
	var trxReq models.ProxyAPITransactionSendWithDataReq
	dataStr := "claimRewards"
	trxReq.Data = base64.StdEncoding.EncodeToString([]byte(dataStr))
	trxReq.GasPrice = 1000000000

	// calculate fee
	totalFee := big.NewInt(0)

	gasForMovement := int64(GAS_FOR_MOVEMENT)
	feeForMovement := big.NewInt(0).Mul(big.NewInt(gasForMovement), big.NewInt(trxReq.GasPrice)) // gasForMovement * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForMovement)

	gasForData := int64(GAS_PER_DATABYTE * len(dataStr))
	feeForData := big.NewInt(0).Mul(big.NewInt(gasForData), big.NewInt(trxReq.GasPrice)) // gasForData * trxReq.GasPrice
	totalFee = big.NewInt(0).Add(totalFee, feeForData)

	feeForExecutionWithoutModifier := big.NewInt(0).Mul(big.NewInt(GAS_FOR_EXECUTION_REWARD), big.NewInt(trxReq.GasPrice)) // gasForExecution * trxReq.GasPrice
	feeForExecution := big.NewInt(0).Div(feeForExecutionWithoutModifier, big.NewInt(100))
	totalFee = big.NewInt(0).Add(totalFee, feeForExecution)

	resp.Value = totalFee.String()

	return
}
