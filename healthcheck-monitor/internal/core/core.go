package core

import (
	"bytes"
	"errors"
	"healthcheck-monitor/internal/config"
	"healthcheck-monitor/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type Core struct {
	Config       *config.Config
	InstanceList []models.InstanceInfo
}

func NewCore(cfg *config.Config) (core *Core, err error) {
	var instanceList []models.InstanceInfo
	for _, v := range cfg.InstanceList {
		var instanceInfo models.InstanceInfo
		instanceInfo.Name = v.Name
		instanceInfo.Endpoint = v.Endpoint
		instanceInfo.Addr, err = getAddrFromEndpoint(v.Endpoint)
		if err != nil {

			return
		}
		instanceList = append(instanceList, instanceInfo)
	}

	core = &Core{
		Config:       cfg,
		InstanceList: instanceList,
	}
	rand.Seed(time.Now().UnixNano())
	go core.ServiceScheduler()

	return
}

func (c *Core) GetServiceAlive() (resp bool, err error) {

	return true, nil
}

func (c *Core) CheckAlive(endpoint string) (isAlive bool) {
	client := http.Client{
		Timeout: 15 * time.Second,
	}

	resHTTP, err := client.Get(endpoint)
	if err != nil {

		return false
	}
	defer resHTTP.Body.Close()
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return false
	}

	if string(body) != "true" {

		return false
	}

	return true
}

func (c *Core) SendBotMessage(messageBody string) {
	jsonData := `
	{
		"chat_id": ` + c.Config.Bot.ChatID + `,
		"text": "#` + c.Config.Bot.Name + ` | ` + c.Config.Monitor.Addr + "\n" + messageBody + `",
		"parse_mode": "HTML"
	}`

	resHTTP, err := http.Post("https://api.telegram.org/bot"+c.Config.Bot.Token+"/sendMessage", "application/json", bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		log.Println(err)
		return
	}

	_, err = ioutil.ReadAll(resHTTP.Body)
	if err != nil {
		log.Println(err)
		return
	}
}

func getAddrFromEndpoint(endpoint string) (addr string, err error) {
	arr1 := strings.Split(endpoint, "//")
	if len(arr1) < 2 {
		err = errors.New("endpoint format is wrong")
		return
	}
	arr2 := strings.Split(arr1[1], ":")
	if len(arr2) == 0 {
		err = errors.New("endpoint format is wrong")
		return
	}
	return arr2[0], nil
}
