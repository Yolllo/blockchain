package models

type InstanceInfo struct {
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
	Addr     string `json:"addr"`
	Status   bool   `json:"status"`
}
