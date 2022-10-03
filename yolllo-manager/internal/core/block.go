package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"yolllo-manager/models"
)

func (c *Core) GetBlockByNonce(req models.GetBlockByNonceReq) (resp models.GetBlockByNonceResp, err error) {
	shardStr := strconv.FormatInt(req.Shard, 10)
	nonceStr := strconv.FormatInt(req.Nonce, 10)
	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/block/" + shardStr + "/by-nonce/" + nonceStr + "?withTxs=true")
	if err != nil {

		return
	}
	defer resHTTP.Body.Close()
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

func (c *Core) GetBlockByHash(req models.GetBlockByHashReq) (resp models.GetBlockByHashResp, err error) {
	shardStr := strconv.FormatInt(req.Shard, 10)
	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/block/" + shardStr + "/by-hash/" + req.Hash + "?withTxs=true")
	if err != nil {

		return
	}
	defer resHTTP.Body.Close()
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

func (c *Core) GetLastBlock(req models.GetLastBlockReq) (resp models.GetLastBlockResp, err error) {
	shardStr := strconv.FormatInt(req.Shard, 10)
	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/network/status/" + shardStr)
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

	resp.Nonce = shardInfo.Data.Status.Nonce

	return
}

func (c *Core) GetLastBlockList(req models.GetLastBlockListReq) (resp models.GetLastBlockListResp, err error) {
	blocks, err := c.Repo.ES.GetLastBlocks(req.PageSize)
	for _, block := range blocks.Hits.Hits {
		var blockInfo models.BlockListBlockInfo
		blockInfo.Hash = block.ID
		blockInfo.Nonce = block.Source.Nonce
		blockInfo.Timestamp = block.Source.Timestamp
		blockInfo.ShardID = block.Source.ShardID
		blockInfo.TxCount = block.Source.TxCount
		resp.BlockList = append(resp.BlockList, blockInfo)
		if len(block.Sort) > 0 {
			resp.NextPageOffset = block.Sort[0]
		}
	}

	return
}

func (c *Core) GetNextBlockList(req models.GetNextBlockListReq) (resp models.GetNextBlockListResp, err error) {
	blocks, err := c.Repo.ES.GetNextBlocks(req.PageSize, req.NextPageOffset)
	for _, block := range blocks.Hits.Hits {
		var blockInfo models.BlockListBlockInfo
		blockInfo.Hash = block.ID
		blockInfo.Nonce = block.Source.Nonce
		blockInfo.Timestamp = block.Source.Timestamp
		blockInfo.ShardID = block.Source.ShardID
		blockInfo.TxCount = block.Source.TxCount
		resp.BlockList = append(resp.BlockList, blockInfo)
		if len(block.Sort) > 0 {
			resp.NextPageOffset = block.Sort[0]
		}
	}

	return
}
