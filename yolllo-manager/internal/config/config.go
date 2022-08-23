package config

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	MnemonicHash string `json:"mnemonic_hash"`
	AuthToken    string `json:"auth_token"`
	ProxyAddress string `json:"proxy_address"`

	Router struct {
		Port string `json:"port"`
	} `json:"router"`

	Swagger struct {
		Host     string `json:"host"`
		IsEnable bool   `json:"is_enable"`
	} `json:"swagger"`

	Repo struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Name     string `json:"name"`
	} `json:"repo"`

	PrivateConfig
}

type PrivateConfig struct {
	Mnemonic string
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

func (cfg *Config) EnterMnemonic() (err error) {
	fmt.Println("Enter a mnemonic to start the service:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	cfg.Mnemonic = scanner.Text()

	inputHashArray := sha256.Sum256([]byte(cfg.Mnemonic))
	inputHash := hex.EncodeToString(inputHashArray[:])
	if inputHash != cfg.MnemonicHash {
		err = errors.New("incorrect mnemonic string (hash=" + inputHash + ")")

		return
	}

	return
}
