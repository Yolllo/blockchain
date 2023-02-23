package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"healthcheck-service/internal/config"
	"healthcheck-service/models"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Core struct {
	Config     *config.Config
	NodeStatus *models.NodeStatus
}

func NewCore(cfg *config.Config) (core *Core, err error) {
	core = &Core{
		Config:     cfg,
		NodeStatus: &models.NodeStatus{},
	}
	rand.Seed(time.Now().UnixNano())
	go core.HeathcheckScheduler()

	return
}

func (c *Core) GetNodeStatus() (resp models.GetNodeStatusResp, err error) {
	resp = models.GetNodeStatusResp{
		ShardID:      c.NodeStatus.ShardID,
		CurrentRound: c.NodeStatus.CurrentRound,
		CurrentEpoch: c.NodeStatus.CurrentEpoch,
	}

	return
}

func (c *Core) GetServiceAlive() (resp bool, err error) {

	return true, nil
}

func (c *Core) UpdateNodeStatus() (err error) {
	urlStr := "http://" + c.Config.NodeAPIAddr + "/node/status"
	resHTTP, err := http.Get(urlStr)
	if err != nil {

		return
	}
	defer resHTTP.Body.Close()
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}

	var resp models.GetNodeMetricsResp
	err = json.Unmarshal(body, &resp)
	if err != nil {

		return
	}

	if resp.Error != "" {
		err = errors.New(resp.Error)

		return
	}

	c.NodeStatus.ShardID = resp.Data.Metrics.ShardID
	c.NodeStatus.CurrentRound = resp.Data.Metrics.CurrentRound
	c.NodeStatus.CurrentEpoch = resp.Data.Metrics.EpochNumber

	return nil
}

func (c *Core) SendBotMessage(messageBody string) {
	jsonData := `
	{
		"chat_id": ` + c.Config.Bot.ChatID + `,
		"text": "#` + c.Config.Bot.Name + ` | ` + c.Config.ServiceAddr + "\n" + messageBody + `",
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

func (c *Core) ExecNode() (err error) {
	log.Println("INIT RESTART")
	os.Chdir(c.Config.NodePath)
	cmd := exec.Command("./start.sh")
	cmd.Run()

	return
}
