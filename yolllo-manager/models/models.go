package models

import "math/big"

type Wallet struct {
	WalletIndex   int64
	WalletAddress string
	CreatedAt     int64
}

type CreateUserAddressResp struct {
	WalletAddress string `json:"wallet_address"`
}

type GetAddressReq struct {
	WalletAddress string `json:"wallet_address"`
}

type GetAddressResp struct {
	Data struct {
		Account struct {
			Address  string `json:"address"`
			Nonce    int64  `json:"nonce"`
			Balanace string `json:"balance"`
		} `json:"account"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type CreateUserTransactionReq struct {
	SenderAddress   string `json:"sender_address"`
	ReceiverAddress string `json:"receiver_address"`
	Value           string `json:"value"`
}

type CreateUserTransactionResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type ProxyAPITransactionSendReq struct {
	Nonce     int64  `json:"nonce"`
	Value     string `json:"value"`
	Receiver  string `json:"receiver"`
	Sender    string `json:"sender"`
	GasPrice  int64  `json:"gasPrice"`
	GasLimit  int64  `json:"gasLimit"`
	Signature string `json:"signature"`
	ChainID   string `json:"chainID"`
	Version   int64  `json:"version"`
}

type ProxyAPITransactionSendWithDataReq struct {
	Nonce     int64  `json:"nonce"`
	Value     string `json:"value"`
	Receiver  string `json:"receiver"`
	Sender    string `json:"sender"`
	GasPrice  int64  `json:"gasPrice"`
	GasLimit  int64  `json:"gasLimit"`
	Data      string `json:"data"`
	Signature string `json:"signature"`
	ChainID   string `json:"chainID"`
	Version   int64  `json:"version"`
}

type ProxyAPITransactionSendResp struct {
	Data struct {
		TxHash string `json:"txHash"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type GetTransactionReq struct {
	TransactionHash string `json:"transaction_hash"`
}

type GetTransactionResp struct {
	Hash          string `json:"hash"`
	Nonce         int64  `json:"nonce"`
	Receiver      string `json:"receiver"`
	Sender        string `json:"sender"`
	ReceiverShard int64  `json:"receiverShard"`
	SenderShard   int64  `json:"senderShard"`
	Value         string `json:"value"`
	Timestamp     int64  `json:"timestamp"`
	Status        string `json:"status"`
}

type GetBlockByNonceReq struct {
	Nonce int64 `json:"nonce"`
	Shard int64 `json:"shard"`
}

type GetBlockByNonceResp struct {
	Data struct {
		Block struct {
			Nonce         int64  `json:"nonce"`
			Round         int64  `json:"round"`
			Hash          string `json:"hash"`
			PrevBlockHash string `json:"prevBlockHash"`
			Epoch         int64  `json:"epoch"`
			Shard         int64  `json:"shard"`
			NumTxs        int64  `json:"numTxs"`
			MiniBlocks    []struct {
				Hash             string `json:"hash"`
				Type             string `json:"type"`
				SourceShard      int64  `json:"sourceShard"`
				DestinationShard int64  `json:"destinationShard"`
				Transactions     []struct {
					Type             string `json:"type"`
					Hash             string `json:"hash"`
					Nonce            int64  `json:"nonce"`
					Value            string `json:"value"`
					Receiver         string `json:"receiver"`
					Sender           string `json:"sender"`
					GasPrice         int64  `json:"gasPrice"`
					GasLimit         int64  `json:"gasLimit"`
					Signature        string `json:"signature"`
					SourceShard      int64  `json:"sourceShard"`
					DestinationShard int64  `json:"destinationShard"`
					MiniBlockType    string `json:"miniblockType"`
					MiniBlockHash    string `json:"miniblockHash"`
					Status           string `json:"status"`
				} `json:"transactions"`
			} `json:"miniBlocks"`
			Timestamp       int64  `json:"timestamp"`
			AccumulatedFees string `json:"accumulatedFees"`
			DeveloperFees   string `json:"developerFees"`
			Status          string `json:"status"`
		} `json:"block"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type GetBlockByHashReq struct {
	Hash  string `json:"hash"`
	Shard int64  `json:"shard"`
}

type GetBlockByHashResp struct {
	Data struct {
		Block struct {
			Nonce         int64  `json:"nonce"`
			Round         int64  `json:"round"`
			Hash          string `json:"hash"`
			PrevBlockHash string `json:"prevBlockHash"`
			Epoch         int64  `json:"epoch"`
			Shard         int64  `json:"shard"`
			NumTxs        int64  `json:"numTxs"`
			MiniBlocks    []struct {
				Hash             string `json:"hash"`
				Type             string `json:"type"`
				SourceShard      int64  `json:"sourceShard"`
				DestinationShard int64  `json:"destinationShard"`
				Transactions     []struct {
					Type             string `json:"type"`
					Hash             string `json:"hash"`
					Nonce            int64  `json:"nonce"`
					Value            string `json:"value"`
					Receiver         string `json:"receiver"`
					Sender           string `json:"sender"`
					GasPrice         int64  `json:"gasPrice"`
					GasLimit         int64  `json:"gasLimit"`
					Signature        string `json:"signature"`
					SourceShard      int64  `json:"sourceShard"`
					DestinationShard int64  `json:"destinationShard"`
					MiniBlockType    string `json:"miniblockType"`
					MiniBlockHash    string `json:"miniblockHash"`
					Status           string `json:"status"`
				} `json:"transactions"`
			} `json:"miniBlocks"`
			Timestamp       int64  `json:"timestamp"`
			AccumulatedFees string `json:"accumulatedFees"`
			DeveloperFees   string `json:"developerFees"`
			Status          string `json:"status"`
		} `json:"block"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type GetTransactionCostReq struct {
	SenderAddress   string `json:"sender_address"`
	ReceiverAddress string `json:"receiver_address"`
	Value           string `json:"value"`
}

type GetTransactionCostResp struct {
	Data struct {
		TxGasUnits int64 `json:"txGasUnits"`
		// TODO field: returnMessage, smartContractResults
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type GetTransactionFeeReq struct {
	SenderAddress   string `json:"sender_address"`
	ReceiverAddress string `json:"receiver_address"`
	Value           string `json:"value"`
}

type GetTransactionFeeResp struct {
	Value *big.Int `json:"value" swaggertype:"integer"`
}

type ProxyAPITransactionCostReq struct {
	Value    string `json:"value"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	ChainID  string `json:"chainID"`
	Version  int64  `json:"version"`
}

type ProxyAPITransactionCostResp struct {
	Data struct {
		TxGasUnits int64 `json:"txGasUnits"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type GetLastBlockReq struct {
	Shard int64 `json:"shard"`
}

type GetLastBlockResp struct {
	Nonce int64 `json:"nonce"`
}

type ProxyAPINetworkStatusShardResp struct {
	Data struct {
		Status struct {
			Round int64 `json:"erd_current_round"`
			Nonce int64 `json:"erd_nonce"`
		} `json:"status"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type CreateTransactionReq struct {
	Nonce           int64  `json:"nonce"` // of sender
	Value           string `json:"value"`
	SenderAddress   string `json:"sender_address"`
	ReceiverAddress string `json:"receiver_address"`
	Signature       string `json:"signature"`
}

type CreateTransactionResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type DelegateUserStakingReq struct {
	UserAddress string `json:"user_address"`
	Value       string `json:"value"`
}

type DelegateUserStakingResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type GetUserStakingRewardReq struct {
	UserAddress string `json:"user_address"`
}

type GetUserStakingRewardResp struct {
	RewardValue string `json:"reward_value"`
}

type ProxyAPIQueryClaimableRewardsReq struct {
	SCAddress string   `json:"scAddress"`
	FuncName  string   `json:"funcName"`
	Args      []string `json:"args"`
}

type ProxyAPIQueryClaimableRewardsResp struct {
	Data struct {
		Data struct {
			ReturnData []string `json:"returnData"`
		} `json:"data"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type ClaimUserStakingRewardReq struct {
	UserAddress string `json:"user_address"`
}

type ClaimUserStakingRewardResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type ProxyAPIQueryGetTotalActiveStakeReq struct {
	SCAddress string   `json:"scAddress"`
	FuncName  string   `json:"funcName"`
	Args      []string `json:"args"`
}

type ProxyAPIQueryGetTotalActiveStakeResp struct {
	Data struct {
		Data struct {
			ReturnData []string `json:"returnData"`
		} `json:"data"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type GetUserStakingTotalStakeResp struct {
	TotalStakeValue string `json:"total_stake_value"`
}

type ElasticAPIQueryGetLastBlocksResp struct {
	Hits struct {
		Hits []struct {
			ID     string `json:"_id"`
			Source struct {
				Nonce       int64 `json:"nonce"`
				Timestamp   int64 `json:"timestamp"`
				ShardID     int64 `json:"shardId"`
				TxCount     int64 `json:"txCount"`
				SearchOrder int64 `json:"searchOrder"`
			} `json:"_source"`
			Sort []float64 `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
}

type BlockListBlockInfo struct {
	Hash      string `json:"hash"`
	Nonce     int64  `json:"nonce"`
	Timestamp int64  `json:"timestamp"`
	ShardID   int64  `json:"shard_id"`
	TxCount   int64  `json:"tx_count"`
}

type GetLastBlockListResp struct {
	NextPageOffset float64              `json:"next_page_offset"`
	BlockList      []BlockListBlockInfo `json:"block_list"`
}

type GetLastBlockListReq struct {
	PageSize int64 `json:"page_size"`
}

type GetNextBlockListReq struct {
	NextPageOffset float64 `json:"next_page_offset"`
	PageSize       int64   `json:"page_size"`
}

type GetNextBlockListResp struct {
	NextPageOffset float64              `json:"next_page_offset"`
	BlockList      []BlockListBlockInfo `json:"block_list"`
}

type ElasticAPIQueryGetLastTransactionsResp struct {
	Hits struct {
		Hits []struct {
			ID     string `json:"_id"`
			Source struct {
				Nonce         int64  `json:"nonce"`
				Receiver      string `json:"receiver"`
				Sender        string `json:"sender"`
				ReceiverShard int64  `json:"receiverShard"`
				SenderShard   int64  `json:"senderShard"`
				Value         string `json:"value"`
				Timestamp     int64  `json:"timestamp"`
				SearchOrder   int64  `json:"searchOrder"`
				Status        string `json:"status"`
			} `json:"_source"`
			Sort []float64 `json:"sort"`
		} `json:"hits"`
		Total struct {
			Value int64 `json:"value"`
		} `json:"total"`
	} `json:"hits"`
}

type TransactionListTransactionInfo struct {
	Hash          string `json:"hash"`
	Nonce         int64  `json:"nonce"`
	Receiver      string `json:"receiver"`
	Sender        string `json:"sender"`
	ReceiverShard int64  `json:"receiverShard"`
	SenderShard   int64  `json:"senderShard"`
	Value         string `json:"value"`
	Timestamp     int64  `json:"timestamp"`
	Status        string `json:"status"`
}

type GetLastTransactionListReq struct {
	PageSize int64 `json:"page_size"`
}

type GetLastTransactionListResp struct {
	NextTimestampAfter   float64                          `json:"next_timestamp_after"`
	NextSearchOrderAfter float64                          `json:"next_searchorder_after"`
	TransactionList      []TransactionListTransactionInfo `json:"transaction_list"`
}

type GetNextTransactionListReq struct {
	PageSize         int64   `json:"page_size"`
	TimestampAfter   float64 `json:"timestamp_after"`
	SearchOrderAfter float64 `json:"searchorder_after"`
}

type GetNextTransactionListResp struct {
	NextTimestampAfter   float64                          `json:"next_timestamp_after"`
	NextSearchOrderAfter float64                          `json:"next_searchorder_after"`
	TransactionList      []TransactionListTransactionInfo `json:"transaction_list"`
}

type GetLastTransactionListByAddrReq struct {
	PageSize      int64  `json:"page_size"`
	WalletAddress string `json:"wallet_address"`
}

type GetLastTransactionListByAddrResp struct {
	NextTimestampAfter   float64                          `json:"next_timestamp_after"`
	NextSearchOrderAfter float64                          `json:"next_searchorder_after"`
	TransactionList      []TransactionListTransactionInfo `json:"transaction_list"`
}

type GetNextTransactionListByAddrReq struct {
	PageSize         int64   `json:"page_size"`
	WalletAddress    string  `json:"wallet_address"`
	TimestampAfter   float64 `json:"timestamp_after"`
	SearchOrderAfter float64 `json:"searchorder_after"`
}

type GetNextTransactionListByAddrResp struct {
	NextTimestampAfter   float64                          `json:"next_timestamp_after"`
	NextSearchOrderAfter float64                          `json:"next_searchorder_after"`
	TransactionList      []TransactionListTransactionInfo `json:"transaction_list"`
}

type GetStakingCurrentMonthlyRewardResp struct {
	Value *big.Int `json:"value" swaggertype:"integer"`
}

type ElasticAPIQueryGetTransactionResp struct {
	Hits struct {
		Hits []struct {
			ID     string `json:"_id"`
			Source struct {
				Nonce         int64  `json:"nonce"`
				Receiver      string `json:"receiver"`
				Sender        string `json:"sender"`
				ReceiverShard int64  `json:"receiverShard"`
				SenderShard   int64  `json:"senderShard"`
				Value         string `json:"value"`
				Timestamp     int64  `json:"timestamp"`
				SearchOrder   int64  `json:"searchOrder"`
				Status        string `json:"status"`
			} `json:"_source"`
			Sort []float64 `json:"sort"`
		} `json:"hits"`
	} `json:"hits"`
}

type GetServerTimeResp struct {
	Timestamp int64 `json:"timestamp"`
}

type GetRangeTransactionListReq struct {
	PageSize      int64 `json:"page_size"`
	PageFrom      int64 `json:"page_from"`
	TimestampFrom int64 `json:"timestamp_from"`
	TimestampTo   int64 `json:"timestamp_to"`
}

type GetRangeTransactionListResp struct {
	Total           int64                            `json:"total"`
	TransactionList []TransactionListTransactionInfo `json:"transaction_list"`
}

type GetUserStakingReq struct {
	UserAddress string `json:"user_address"`
}

type GetUserStakingResp struct {
	Value string `json:"value"`
}

type ProxyAPIQueryDelegatorActiveStakeReq struct {
	SCAddress string   `json:"scAddress"`
	FuncName  string   `json:"funcName"`
	Args      []string `json:"args"`
}

type UndelegateUserStakingReq struct {
	UserAddress string `json:"user_address"`
	Value       string `json:"value"`
}

type UndelegateUserStakingResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type GetUserStakingUndelegatedReq struct {
	UserAddress string `json:"user_address"`
}

type GetUserStakingUndelegatedResp struct {
	Values []string `json:"value"`
}

type ProxyAPIQueryGetUndelegatedListResp struct {
	Data struct {
		Data struct {
			ReturnData []string `json:"returnData"`
		} `json:"data"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type ClaimUserStakingUndelegatedReq struct {
	UserAddress string `json:"user_address"`
}

type ClaimUserStakingUndelegatedResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type GetUserStakingFeeResp struct {
	Value float64 `json:"value"`
}

type ProxyAPIQueryGetContractConfigReq struct {
	SCAddress string `json:"scAddress"`
	FuncName  string `json:"funcName"`
}

type ProxyAPIQueryGetContractConfigResp struct {
	Data struct {
		Data struct {
			ReturnData []string `json:"returnData"`
		} `json:"data"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type SetUserStakingFeeReq struct {
	Value float64 `json:"value"`
}

type SetUserStakingFeeResp struct {
	TransactionHash string `json:"transaction_hash"`
}

type IsValidAddressReq struct {
	WalletAddress string `json:"wallet_address"`
}

type IsValidAddressResp struct {
	IsValid bool `json:"is_valid"`
}
