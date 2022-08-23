package router

import (
	"net/http"
	"yolllo-manager/models"

	"github.com/gin-gonic/gin"
)

// CreateUserAddress godoc
// @Summary create user address
// @Schemes
// @Param Authorization header string true "Authorization"
// @Description create yol-address for user
// @Tags address
// @Accept json
// @Produce json
// @Success 200 {object} models.CreateUserAddressResp
// @Failure 400
// @Router /user/address/create [post]
func (r *Router) CreateUserAddress(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) != 1 {
		c.String(http.StatusBadRequest, "missing auth token")
		return
	}

	authToken := c.Request.Header["Authorization"][0]
	if authToken != r.Config.AuthToken {
		c.String(http.StatusBadRequest, "wrong auth token")
		return
	}

	resp, err := r.Core.CreateUserAddress()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateUserTransaction godoc
// @Summary create user transaction
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.CreateUserTransactionReq true "CreateTransactionUserReq params"
// @Description create transaction for user
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.CreateUserTransactionResp
// @Failure 400
// @Router /user/transaction/create [post]
func (r *Router) CreateUserTransaction(c *gin.Context) {
	var req models.CreateUserTransactionReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if len(c.Request.Header["Authorization"]) != 1 {
		c.String(http.StatusBadRequest, "missing auth token")
		return
	}

	authToken := c.Request.Header["Authorization"][0]
	if authToken != r.Config.AuthToken {
		c.String(http.StatusBadRequest, "wrong auth token")
		return
	}

	resp, err := r.Core.CreateUserTransaction(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetAddress godoc
// @Summary get address
// @Schemes
// @Param JSON body models.GetAddressReq true "GetAddressReq params"
// @Description get address full info
// @Tags address
// @Accept json
// @Produce json
// @Success 200 {object} models.GetAddressResp
// @Failure 400
// @Router /address/get [post]
func (r *Router) GetAddress(c *gin.Context) {
	var req models.GetAddressReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetWalletBalance(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetTransaction godoc
// @Summary get transaction
// @Schemes
// @Param JSON body models.GetTransactionReq true "GetTransactionReq params"
// @Description get transaction full info
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetTransactionResp
// @Failure 400
// @Router /transaction/get [post]
func (r *Router) GetTransaction(c *gin.Context) {
	var req models.GetTransactionReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetTransaction(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetBlockByNonce godoc
// @Summary get block
// @Schemes
// @Param JSON body models.GetBlockByNonceReq true "GetBlockByNonceReq params"
// @Description get block full info
// @Tags block
// @Accept json
// @Produce json
// @Success 200 {object} models.GetBlockByNonceResp
// @Failure 400
// @Router /block/by-nonce/get [post]
func (r *Router) GetBlockByNonce(c *gin.Context) {
	var req models.GetBlockByNonceReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetBlockByNonce(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetBlockByHash godoc
// @Summary get block
// @Schemes
// @Param JSON body models.GetBlockByHashReq true "GetBlockByHashReq params"
// @Description get block full info
// @Tags block
// @Accept json
// @Produce json
// @Success 200 {object} models.GetBlockByHashResp
// @Failure 400
// @Router /block/by-hash/get [post]
func (r *Router) GetBlockByHash(c *gin.Context) {
	var req models.GetBlockByHashReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetBlockByHash(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetTransactionCost godoc
// @Summary transaction cost
// @Schemes
// @Param JSON body models.GetTransactionCostReq true "GetTransactionCostReq params"
// @Description get transaction cost
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetTransactionCostResp
// @Failure 400
// @Router /transaction/cost [post]
func (r *Router) GetTransactionCost(c *gin.Context) {
	var req models.GetTransactionCostReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetTransactionCost(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLastBlock godoc
// @Summary last block
// @Schemes
// @Param JSON body models.GetLastBlockReq true "GetLastBlockReq params"
// @Description get last nonce in shard
// @Tags block
// @Accept json
// @Produce json
// @Success 200 {object} models.GetLastBlockResp
// @Failure 400
// @Router /block/last [post]
func (r *Router) GetLastBlock(c *gin.Context) {
	var req models.GetLastBlockReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetLastBlock(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// CreateTransaction godoc
// @Summary create transaction
// @Schemes
// @Param JSON body models.CreateTransactionReq true "CreateTransactionReq params"
// @Description create transaction
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.CreateTransactionResp
// @Failure 400
// @Router /transaction/create [post]
func (r *Router) CreateTransaction(c *gin.Context) {
	var req models.CreateTransactionReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.CreateTransaction(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
