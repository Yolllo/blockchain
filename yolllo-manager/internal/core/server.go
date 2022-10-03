package core

import (
	"time"
	"yolllo-manager/models"
)

func (c *Core) GetServerTime() (resp models.GetServerTimeResp, err error) {
	resp.Timestamp = time.Now().Unix()

	return
}
