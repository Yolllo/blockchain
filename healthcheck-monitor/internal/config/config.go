package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Monitor struct {
		Addr            string `json:"addr"`
		Port            string `json:"port"`
		RedundancyLevel int    `json:"redundancy_level"`
		SecondAddr      string `json:"second_addr"`
	} `json:"monitor"`
	Bot struct {
		Name   string `json:"name"`
		Token  string `json:"token"`
		ChatID string `json:"chat_id"`
	} `json:"bot"`
	InstanceList []struct {
		Name     string `json:"name"`
		Endpoint string `json:"endpoint"`
	} `json:"instance_list"`
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
