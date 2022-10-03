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
	MnemonicHash   string `json:"mnemonic_hash"`
	AuthToken      string `json:"auth_token"`
	ProxyAddress   string `json:"proxy_address"`
	StakingAddress string `json:"staking_address"`
	StakingOwner   string `json:"staking_owner"`

	Router struct {
		Port string `json:"port"`
	} `json:"router"`

	Swagger struct {
		Host     string `json:"host"`
		IsEnable bool   `json:"is_enable"`
	} `json:"swagger"`

	Repo struct {
		PostgreSQL struct {
			Host         string `json:"host"`
			Port         string `json:"port"`
			User         string `json:"user"`
			PasswordHash string `json:"password_hash"`
			Name         string `json:"name"`
		} `json:"PostgreSQL"`
		ElasticSearch struct {
			Host         string `json:"host"`
			Port         string `json:"port"`
			User         string `json:"user"`
			PasswordHash string `json:"password_hash"`
		} `json:"ElasticSearch"`
	} `json:"repo"`

	PrivateConfig
}

type PrivateConfig struct {
	Mnemonic              string
	PostgreSQLPassword    string
	ElasticSearchPassword string
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

func (cfg *Config) EnterInitData() (err error) {
	scanner := bufio.NewScanner(os.Stdin)

	// mnemonic
	fmt.Println("Enter a Mnemonic string:")
	scanner.Scan()
	cfg.Mnemonic = scanner.Text()
	inputHashArray := sha256.Sum256([]byte(cfg.Mnemonic))
	inputHash := hex.EncodeToString(inputHashArray[:])
	if inputHash != cfg.MnemonicHash {
		err = errors.New("incorrect mnemonic string (hash=" + inputHash + ")")

		return
	}

	// postgresql
	fmt.Println("Enter PostgreSQL password:")
	scanner.Scan()
	cfg.PostgreSQLPassword = scanner.Text()
	inputHashArray = sha256.Sum256([]byte(cfg.PostgreSQLPassword))
	inputHash = hex.EncodeToString(inputHashArray[:])
	if inputHash != cfg.Repo.PostgreSQL.PasswordHash {
		err = errors.New("incorrect postgresql password (hash=" + inputHash + ")")

		return
	}

	fmt.Println("Enter ElasticSearch password:")
	scanner.Scan()
	cfg.ElasticSearchPassword = scanner.Text()
	inputHashArray = sha256.Sum256([]byte(cfg.ElasticSearchPassword))
	inputHash = hex.EncodeToString(inputHashArray[:])
	if inputHash != cfg.Repo.ElasticSearch.PasswordHash {
		err = errors.New("incorrect elasticsearch password (hash=" + inputHash + ")")

		return
	}

	return
}
