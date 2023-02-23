package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ServiceAddr string `json:"service_addr"`
	ServicePort string `json:"service_port"`
	NodeAPIAddr string `json:"node_api_addr"`
	NodePath    string `json:"node_path"`
	Bot         struct {
		Name   string `json:"name"`
		Token  string `json:"token"`
		ChatID string `json:"chat_id"`
	} `json:"bot"`
}

func LoadConfig() (cfg *Config, err error) {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {

		return
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {

		return
	}

	return
}
