package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"yolllo-manager/models"
)

func (c *Core) CreateUserAddress() (resp models.CreateUserAddressResp, err error) {
	walletAddress, err := c.Repo.PG.CreateNewWallet()
	if err != nil {

		return
	}

	resp.WalletAddress = walletAddress

	return
}

func (c *Core) GetWalletBalance(req models.GetAddressReq) (resp models.GetAddressResp, err error) {
	resHTTP, err := http.Get("http://" + c.Config.ProxyAddress + "/address/" + req.WalletAddress)
	if err != nil {

		return
	}
	defer resHTTP.Body.Close()
	body, err := ioutil.ReadAll(resHTTP.Body)
	if err != nil {

		return
	}
	err = json.Unmarshal(body, &resp)
	if err != nil {

		return
	}

	if resp.Error != "" {
		err = errors.New(resp.Error)

		return
	}

	return
}
