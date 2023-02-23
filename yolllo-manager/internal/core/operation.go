package core

import (
	"encoding/base64"
	"yolllo-manager/models"
)

func (c *Core) GetLastOperationList(req models.GetLastOperationListReq) (resp models.GetLastOperationListResp, err error) {
	operations, err := c.Repo.ES.GetLastOperations(req.PageSize)
	if err != nil {

		return
	}
	for _, operation := range operations.Hits.Hits {
		var operationInfo models.OperationListOperationInfo
		operationInfo.Hash = operation.ID
		operationInfo.Nonce = operation.Source.Nonce
		operationInfo.Receiver = operation.Source.Receiver
		operationInfo.Sender = operation.Source.Sender
		operationInfo.ReceiverShard = operation.Source.ReceiverShard
		operationInfo.SenderShard = operation.Source.SenderShard
		operationInfo.Value = operation.Source.Value
		operationInfo.Timestamp = operation.Source.Timestamp
		operationInfo.Status = operation.Source.Status
		operationInfo.Operation = operation.Source.Operation
		operationInfo.Fee = operation.Source.Fee
		operationInfo.Function = operation.Source.Function
		if len(operation.Source.Data) > 0 {
			dataByte, err := base64.StdEncoding.DecodeString(operation.Source.Data)
			if err != nil {

				return resp, err
			}
			operationInfo.Data = string(dataByte)
		}
		resp.OperationList = append(resp.OperationList, operationInfo)
		if len(operation.Sort) > 0 {
			resp.NextTimestampAfter = operation.Sort[0]
			resp.NextSearchOrderAfter = operation.Sort[1]
		}
	}

	return
}

func (c *Core) GetNextOperationList(req models.GetNextOperationListReq) (resp models.GetNextOperationListResp, err error) {
	operations, err := c.Repo.ES.GetNextOperations(req.PageSize, req.TimestampAfter, req.SearchOrderAfter)
	if err != nil {

		return
	}
	for _, operation := range operations.Hits.Hits {
		var operationInfo models.OperationListOperationInfo
		operationInfo.Hash = operation.ID
		operationInfo.Nonce = operation.Source.Nonce
		operationInfo.Receiver = operation.Source.Receiver
		operationInfo.Sender = operation.Source.Sender
		operationInfo.ReceiverShard = operation.Source.ReceiverShard
		operationInfo.SenderShard = operation.Source.SenderShard
		operationInfo.Value = operation.Source.Value
		operationInfo.Timestamp = operation.Source.Timestamp
		operationInfo.Status = operation.Source.Status
		operationInfo.Operation = operation.Source.Operation
		operationInfo.Fee = operation.Source.Fee
		operationInfo.Function = operation.Source.Function
		if len(operation.Source.Data) > 0 {
			dataByte, err := base64.StdEncoding.DecodeString(operation.Source.Data)
			if err != nil {

				return resp, err
			}
			operationInfo.Data = string(dataByte)
		}
		resp.OperationList = append(resp.OperationList, operationInfo)
		if len(operation.Sort) > 0 {
			resp.NextTimestampAfter = operation.Sort[0]
			resp.NextSearchOrderAfter = operation.Sort[1]
		}
	}

	return
}

func (c *Core) GetRangeOperationList(req models.GetRangeOperationListReq) (resp models.GetRangeOperationListResp, err error) {
	operations, err := c.Repo.ES.GetRangeOperations(req.PageSize, req.PageFrom, req.TimestampFrom, req.TimestampTo)
	if err != nil {

		return
	}
	for _, operation := range operations.Hits.Hits {
		var operationInfo models.OperationListOperationInfo
		operationInfo.Hash = operation.ID
		operationInfo.Nonce = operation.Source.Nonce
		operationInfo.Receiver = operation.Source.Receiver
		operationInfo.Sender = operation.Source.Sender
		operationInfo.ReceiverShard = operation.Source.ReceiverShard
		operationInfo.SenderShard = operation.Source.SenderShard
		operationInfo.Value = operation.Source.Value
		operationInfo.Timestamp = operation.Source.Timestamp
		operationInfo.Status = operation.Source.Status
		operationInfo.Operation = operation.Source.Operation
		operationInfo.Fee = operation.Source.Fee
		operationInfo.Function = operation.Source.Function
		if len(operation.Source.Data) > 0 {
			dataByte, err := base64.StdEncoding.DecodeString(operation.Source.Data)
			if err != nil {

				return resp, err
			}
			operationInfo.Data = string(dataByte)
		}
		resp.OperationList = append(resp.OperationList, operationInfo)
	}
	resp.Total = operations.Hits.Total.Value

	return
}

func (c *Core) GetLastOperationListByAddr(req models.GetLastOperationListByAddrReq) (resp models.GetLastOperationListByAddrResp, err error) {
	operations, err := c.Repo.ES.GetLastOperationsByAddr(req.PageSize, req.WalletAddress)
	if err != nil {

		return
	}
	for _, operation := range operations.Hits.Hits {
		var operationInfo models.OperationListOperationInfo
		operationInfo.Hash = operation.ID
		operationInfo.Nonce = operation.Source.Nonce
		operationInfo.Receiver = operation.Source.Receiver
		operationInfo.Sender = operation.Source.Sender
		operationInfo.ReceiverShard = operation.Source.ReceiverShard
		operationInfo.SenderShard = operation.Source.SenderShard
		operationInfo.Value = operation.Source.Value
		operationInfo.Timestamp = operation.Source.Timestamp
		operationInfo.Status = operation.Source.Status
		operationInfo.Operation = operation.Source.Operation
		operationInfo.Fee = operation.Source.Fee
		operationInfo.Function = operation.Source.Function
		if len(operation.Source.Data) > 0 {
			dataByte, err := base64.StdEncoding.DecodeString(operation.Source.Data)
			if err != nil {

				return resp, err
			}
			operationInfo.Data = string(dataByte)
		}
		resp.OperationList = append(resp.OperationList, operationInfo)
		if len(operation.Sort) > 0 {
			resp.NextTimestampAfter = operation.Sort[0]
			resp.NextSearchOrderAfter = operation.Sort[1]
		}
	}

	return
}

func (c *Core) GetNextOperationListByAddr(req models.GetNextOperationListByAddrReq) (resp models.GetNextOperationListByAddrResp, err error) {
	operations, err := c.Repo.ES.GetNextOperationsByAddr(req.PageSize, req.WalletAddress, req.TimestampAfter, req.SearchOrderAfter)
	if err != nil {

		return
	}
	for _, operation := range operations.Hits.Hits {
		var operationInfo models.OperationListOperationInfo
		operationInfo.Hash = operation.ID
		operationInfo.Nonce = operation.Source.Nonce
		operationInfo.Receiver = operation.Source.Receiver
		operationInfo.Sender = operation.Source.Sender
		operationInfo.ReceiverShard = operation.Source.ReceiverShard
		operationInfo.SenderShard = operation.Source.SenderShard
		operationInfo.Value = operation.Source.Value
		operationInfo.Timestamp = operation.Source.Timestamp
		operationInfo.Status = operation.Source.Status
		operationInfo.Operation = operation.Source.Operation
		operationInfo.Fee = operation.Source.Fee
		operationInfo.Function = operation.Source.Function
		if len(operation.Source.Data) > 0 {
			dataByte, err := base64.StdEncoding.DecodeString(operation.Source.Data)
			if err != nil {

				return resp, err
			}
			operationInfo.Data = string(dataByte)
		}
		resp.OperationList = append(resp.OperationList, operationInfo)
		if len(operation.Sort) > 0 {
			resp.NextTimestampAfter = operation.Sort[0]
			resp.NextSearchOrderAfter = operation.Sort[1]
		}
	}

	return
}

func (c *Core) GetOperation(req models.GetOperationReq) (resp models.GetOperationResp, err error) {
	operations, err := c.Repo.ES.GetOperationByHash(req.OperationHash)
	if err != nil {

		return
	}
	for _, operation := range operations.Hits.Hits {
		resp.Hash = operation.ID
		resp.Nonce = operation.Source.Nonce
		resp.Receiver = operation.Source.Receiver
		resp.Sender = operation.Source.Sender
		resp.ReceiverShard = operation.Source.ReceiverShard
		resp.SenderShard = operation.Source.SenderShard
		resp.Value = operation.Source.Value
		resp.Timestamp = operation.Source.Timestamp
		resp.Status = operation.Source.Status
		resp.Operation = operation.Source.Operation
		resp.Fee = operation.Source.Fee
		resp.Function = operation.Source.Function
		if len(operation.Source.Data) > 0 {
			dataByte, err := base64.StdEncoding.DecodeString(operation.Source.Data)
			if err != nil {

				return resp, err
			}
			resp.Data = string(dataByte)
		}

		break
	}

	return
}
