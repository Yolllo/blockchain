package models

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
	Data struct {
		Transaction struct {
			Type             string `json:"type"`
			Nonce            int64  `json:"nonce"`
			Value            string `json:"value"`
			Receiver         string `json:"receiver"`
			Sender           string `json:"sender"`
			GasPrice         int64  `json:"gasPrice"`
			GasLimit         int64  `json:"gasLimit"`
			Signature        string `json:"signature"`
			SourceShard      int64  `json:"sourceShard"`
			DestinationShard int64  `json:"destinationShard"`
			Status           string `json:"status"`
		} `json:"transaction"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
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
		Nonce int64 `json:"txGasUnits"`
		// TODO field: returnMessage, smartContractResults
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}

type ProxyAPITransactionCostReq struct {
	Value    string `json:"value"`
	Receiver string `json:"receiver"`
	Sender   string `json:"sender"`
	ChainID  string `json:"chainID"`
	Version  int64  `json:"version"`
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
