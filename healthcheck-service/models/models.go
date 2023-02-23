package models

type NodeStatus struct {
	InactivityCounter int
	ShardID           int64
	CurrentRound      int64
	CurrentEpoch      int64
}

type GetNodeStatusResp struct {
	IsActive     bool  `json:"is_active"`
	ShardID      int64 `json:"shard_id"`
	CurrentRound int64 `json:"current_round"`
	CurrentEpoch int64 `json:"current_epoch"`
}

type GetNodeMetricsResp struct {
	Data struct {
		Metrics struct {
			ShardID      int64 `json:"erd_shard_id"`
			CurrentRound int64 `json:"erd_current_round"`
			EpochNumber  int64 `json:"erd_epoch_number"`
		} `json:"metrics"`
	} `json:"data"`
	Error string `json:"error"`
	Code  string `json:"code"`
}
