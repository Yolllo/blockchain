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
// @Description get transaction cost in units value
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

// GetTransactionFee godoc
// @Summary transaction fee
// @Schemes
// @Param JSON body models.GetTransactionFeeReq true "GetTransactionFeeReq params"
// @Description get transaction fee
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetTransactionFeeResp
// @Failure 400
// @Router /transaction/fee [post]
func (r *Router) GetTransactionFee(c *gin.Context) {
	var req models.GetTransactionFeeReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetTransactionFee(req)
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

// DelegateUserStaking godoc
// @Summary delegate user staking
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.DelegateUserStakingReq true "DelegateUserStakingReq params"
// @Description delegate to staking, fee ~0,056 YOL
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.DelegateUserStakingResp
// @Failure 400
// @Router /user/staking/delegate [post]
func (r *Router) DelegateUserStaking(c *gin.Context) {
	var req models.DelegateUserStakingReq
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

	resp, err := r.Core.DelegateUserStaking(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DelegateFeeUserStaking godoc
// @Summary delegate user staking fee
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.DelegateUserStakingReq true "DelegateUserStakingReq params"
// @Description check fee for delegate to staking transaction
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.DelegateFeeUserStakingResp
// @Failure 400
// @Router /user/staking/delegate-fee [post]
func (r *Router) DelegateFeeUserStaking(c *gin.Context) {
	var req models.DelegateUserStakingReq
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

	resp, err := r.Core.DelegateFeeUserStaking(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserStaking godoc
// @Summary get user staking
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.GetUserStakingReq true "GetUserStakingReq params"
// @Description get active user stake
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserStakingResp
// @Failure 400
// @Router /user/staking/get [post]
func (r *Router) GetUserStaking(c *gin.Context) {
	var req models.GetUserStakingReq
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

	resp, err := r.Core.GetUserStaking(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UndelegateUserStaking godoc
// @Summary undelegate user staking
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.UndelegateUserStakingReq true "UndelegateUserStakingReq params"
// @Description undelegate from staking, fee ~0,056 YOL
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.UndelegateUserStakingResp
// @Failure 400
// @Router /user/staking/undelegate [post]
func (r *Router) UndelegateUserStaking(c *gin.Context) {
	var req models.UndelegateUserStakingReq
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

	resp, err := r.Core.UndelegateUserStaking(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// UndelegateFeeUserStaking godoc
// @Summary undelegate user staking fee
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.UndelegateUserStakingReq true "UndelegateUserStakingReq params"
// @Description check fee for undelegate from staking transaction, the value may vary depending on the undelegate amount
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.UndelegateFeeUserStakingResp
// @Failure 400
// @Router /user/staking/undelegate-fee [post]
func (r *Router) UndelegateFeeUserStaking(c *gin.Context) {
	var req models.UndelegateUserStakingReq
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

	resp, err := r.Core.UndelegateFeeUserStaking(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserStakingUndelegated godoc
// @Summary get user staking undelegated
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.GetUserStakingUndelegatedReq true "GetUserStakingUndelegatedReq params"
// @Description get user undelegated values that are ready for claim
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserStakingUndelegatedResp
// @Failure 400
// @Router /user/staking/undelegated/get [post]
func (r *Router) GetUserStakingUndelegated(c *gin.Context) {
	var req models.GetUserStakingUndelegatedReq
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

	resp, err := r.Core.GetUserStakingUndelegated(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ClaimUserStakingUndelegated godoc
// @Summary claim user staking undelegated
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.ClaimUserStakingUndelegatedReq true "ClaimUserStakingUndelegatedReq params"
// @Description claim user undelegated coins
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.ClaimUserStakingUndelegatedResp
// @Failure 400
// @Router /user/staking/undelegated/claim [post]
func (r *Router) ClaimUserStakingUndelegated(c *gin.Context) {
	var req models.ClaimUserStakingUndelegatedReq
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

	resp, err := r.Core.ClaimUserStakingUndelegated(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ClaimFeeUserStakingUndelegated godoc
// @Summary claim user staking undelegated fee
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.ClaimUserStakingUndelegatedReq true "ClaimUserStakingUndelegatedReq params"
// @Description check fee for claim user undelegated coins
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.ClaimFeeUserStakingUndelegatedResp
// @Failure 400
// @Router /user/staking/undelegated/claim-fee [post]
func (r *Router) ClaimFeeUserStakingUndelegated(c *gin.Context) {
	var req models.ClaimUserStakingUndelegatedReq
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

	resp, err := r.Core.ClaimFeeUserStakingUndelegated(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserStakingReward godoc
// @Summary get user staking reward
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.GetUserStakingRewardReq true "GetUserStakingRewardReq params"
// @Description get user reward value
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserStakingRewardResp
// @Failure 400
// @Router /user/staking/reward/get [post]
func (r *Router) GetUserStakingReward(c *gin.Context) {
	var req models.GetUserStakingRewardReq
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

	resp, err := r.Core.GetUserStakingReward(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ClaimUserStakingReward godoc
// @Summary claim user staking reward
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.ClaimUserStakingRewardReq true "ClaimUserStakingRewardReq params"
// @Description claim user reward, fee 0,006 YOL
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.ClaimUserStakingRewardResp
// @Failure 400
// @Router /user/staking/reward/claim [post]
func (r *Router) ClaimUserStakingReward(c *gin.Context) {
	var req models.ClaimUserStakingRewardReq
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

	resp, err := r.Core.ClaimUserStakingReward(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ClaimFeeUserStakingReward godoc
// @Summary claim user staking reward fee
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.ClaimUserStakingRewardReq true "ClaimUserStakingRewardReq params"
// @Description check fee for claim user reward
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.ClaimFeeUserStakingRewardResp
// @Failure 400
// @Router /user/staking/reward/claim-fee [post]
func (r *Router) ClaimFeeUserStakingReward(c *gin.Context) {
	var req models.ClaimUserStakingRewardReq
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

	resp, err := r.Core.ClaimFeeUserStakingReward(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// LastClaimedUserStakingReward godoc
// @Summary get last claimed user staking reward
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.LastClaimedUserStakingRewardReq true "LastClaimedUserStakingRewardReq params"
// @Description get last claimed user reward after timestamp of claim trx
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.LastClaimedUserStakingRewardResp
// @Failure 400
// @Router /user/staking/reward/last-claimed [post]
func (r *Router) LastClaimedUserStakingReward(c *gin.Context) {
	var req models.LastClaimedUserStakingRewardReq
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

	resp, err := r.Core.LastClaimedUserStakingReward(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserStakingTotalStake godoc
// @Summary get user total staking
// @Schemes
// @Param Authorization header string true "Authorization"
// @Description get user total stake value
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserStakingTotalStakeResp
// @Failure 400
// @Router /user/staking/total-stake/get [post]
func (r *Router) GetUserStakingTotalStake(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) != 1 {
		c.String(http.StatusBadRequest, "missing auth token")
		return
	}

	authToken := c.Request.Header["Authorization"][0]
	if authToken != r.Config.AuthToken {
		c.String(http.StatusBadRequest, "wrong auth token")
		return
	}

	resp, err := r.Core.GetUserStakingTotalStake()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserStakingTotalReward godoc
// @Summary get user total reward
// @Schemes
// @Param Authorization header string true "Authorization"
// @Description get user total reward value
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserStakingTotalRewardResp
// @Failure 400
// @Router /user/staking/total-reward/get [post]
func (r *Router) GetUserStakingTotalReward(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) != 1 {
		c.String(http.StatusBadRequest, "missing auth token")
		return
	}

	authToken := c.Request.Header["Authorization"][0]
	if authToken != r.Config.AuthToken {
		c.String(http.StatusBadRequest, "wrong auth token")
		return
	}

	resp, err := r.Core.GetUserStakingTotalReward()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLastBlockList godoc
// @Summary last block list
// @Schemes
// @Param JSON body models.GetLastBlockListReq true "GetLastBlockListReq params"
// @Description get last block list page
// @Tags block
// @Accept json
// @Produce json
// @Success 200 {object} models.GetLastBlockListResp
// @Failure 400
// @Router /block/list/last [post]
func (r *Router) GetLastBlockList(c *gin.Context) {
	var req models.GetLastBlockListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetLastBlockList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNextBlockList godoc
// @Summary next block list
// @Schemes
// @Param JSON body models.GetNextBlockListReq true "GetNextBlockListReq params"
// @Description get next block list page
// @Tags block
// @Accept json
// @Produce json
// @Success 200 {object} models.GetNextBlockListResp
// @Failure 400
// @Router /block/list/next [post]
func (r *Router) GetNextBlockList(c *gin.Context) {
	var req models.GetNextBlockListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetNextBlockList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLastTransactionList godoc
// @Summary last transaction list
// @Schemes
// @Param JSON body models.GetLastTransactionListReq true "GetLastTransactionListReq params"
// @Description get last transaction list page
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetLastTransactionListResp
// @Failure 400
// @Router /transaction/list/last [post]
func (r *Router) GetLastTransactionList(c *gin.Context) {
	var req models.GetLastTransactionListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetLastTransactionList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNextTransactionList godoc
// @Summary next transaction list
// @Schemes
// @Param JSON body models.GetNextTransactionListReq true "GetNextTransactionListReq params"
// @Description get next transaction list page
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetNextTransactionListResp
// @Failure 400
// @Router /transaction/list/next [post]
func (r *Router) GetNextTransactionList(c *gin.Context) {
	var req models.GetNextTransactionListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetNextTransactionList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLastTransactionListByAddr godoc
// @Summary last transaction list by address
// @Schemes
// @Param JSON body models.GetLastTransactionListByAddrReq true "GetLastTransactionListByAddrReq params"
// @Description get last transaction list page for address
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetLastTransactionListByAddrResp
// @Failure 400
// @Router /transaction/list/by-address/last [post]
func (r *Router) GetLastTransactionListByAddr(c *gin.Context) {
	var req models.GetLastTransactionListByAddrReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetLastTransactionListByAddr(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNextTransactionListByAddr godoc
// @Summary next transaction list by address
// @Schemes
// @Param JSON body models.GetNextTransactionListByAddrReq true "GetNextTransactionListByAddrReq params"
// @Description get next transaction list page for address
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetNextTransactionListByAddrResp
// @Failure 400
// @Router /transaction/list/by-address/next [post]
func (r *Router) GetNextTransactionListByAddr(c *gin.Context) {
	var req models.GetNextTransactionListByAddrReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetNextTransactionListByAddr(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetStakingCurrentMonthlyReward godoc
// @Summary get staking current monthly reward
// @Schemes
// @Description get reward for current month by staking
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetStakingCurrentMonthlyRewardResp
// @Failure 400
// @Router /staking/current-reward/monthly [post]
func (r *Router) GetStakingCurrentMonthlyReward(c *gin.Context) {
	resp, err := r.Core.GetStakingCurrentMonthlyReward()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetServerTime godoc
// @Summary get server time
// @Schemes
// @Description get server time
// @Tags server
// @Accept json
// @Produce json
// @Success 200 {object} models.GetServerTimeResp
// @Failure 400
// @Router /server/time/get [post]
func (r *Router) GetServerTime(c *gin.Context) {
	resp, err := r.Core.GetServerTime()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRangeTransactionList godoc
// @Summary range transaction list
// @Schemes
// @Param JSON body models.GetRangeTransactionListReq true "GetRangeTransactionListReq params"
// @Description get range transaction list page
// @Tags transaction
// @Accept json
// @Produce json
// @Success 200 {object} models.GetRangeTransactionListResp
// @Failure 400
// @Router /transaction/list/range [post]
func (r *Router) GetRangeTransactionList(c *gin.Context) {
	var req models.GetRangeTransactionListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetRangeTransactionList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetUserStakingFee godoc
// @Summary get user staking fee
// @Schemes
// @Param Authorization header string true "Authorization"
// @Description get user staking fee
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.GetUserStakingFeeResp
// @Failure 400
// @Router /user/staking/fee/get [post]
func (r *Router) GetUserStakingFee(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) != 1 {
		c.String(http.StatusBadRequest, "missing auth token")
		return
	}

	authToken := c.Request.Header["Authorization"][0]
	if authToken != r.Config.AuthToken {
		c.String(http.StatusBadRequest, "wrong auth token")
		return
	}

	resp, err := r.Core.GetUserStakingFee()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// SetUserStakingFee godoc
// @Summary set user staking fee
// @Schemes
// @Param Authorization header string true "Authorization"
// @Param JSON body models.SetUserStakingFeeReq true "SetUserStakingFeeReq params"
// @Description set user staking fee, min 0, max 99.99, step 0.01
// @Tags staking
// @Accept json
// @Produce json
// @Success 200 {object} models.SetUserStakingFeeResp
// @Failure 400
// @Router /user/staking/fee/set [post]
func (r *Router) SetUserStakingFee(c *gin.Context) {
	var req models.SetUserStakingFeeReq
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

	resp, err := r.Core.SetUserStakingFee(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// IsValidAddress godoc
// @Summary check addr for valid
// @Schemes
// @Param JSON body models.IsValidAddressReq true "IsValidAddressReq params"
// @Description check addr for valid
// @Tags address
// @Accept json
// @Produce json
// @Success 200 {object} models.IsValidAddressResp
// @Failure 400
// @Router /address/is-valid [post]
func (r *Router) IsValidAddress(c *gin.Context) {
	var req models.IsValidAddressReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.IsValidAddress(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLastOperationList godoc
// @Summary last operation list
// @Schemes
// @Param JSON body models.GetLastOperationListReq true "GetLastOperationListReq params"
// @Description get last operation list page
// @Tags operation
// @Accept json
// @Produce json
// @Success 200 {object} models.GetLastOperationListResp
// @Failure 400
// @Router /operation/list/last [post]
func (r *Router) GetLastOperationList(c *gin.Context) {
	var req models.GetLastOperationListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetLastOperationList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNextOperationList godoc
// @Summary next operation list
// @Schemes
// @Param JSON body models.GetNextOperationListReq true "GetNextOperationListReq params"
// @Description get next operation list page
// @Tags operation
// @Accept json
// @Produce json
// @Success 200 {object} models.GetNextOperationListResp
// @Failure 400
// @Router /operation/list/next [post]
func (r *Router) GetNextOperationList(c *gin.Context) {
	var req models.GetNextOperationListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetNextOperationList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRangeOperationList godoc
// @Summary range operation list
// @Schemes
// @Param JSON body models.GetRangeOperationListReq true "GetRangeOperationListReq params"
// @Description get range operation list page
// @Tags operation
// @Accept json
// @Produce json
// @Success 200 {object} models.GetRangeOperationListResp
// @Failure 400
// @Router /operation/list/range [post]
func (r *Router) GetRangeOperationList(c *gin.Context) {
	var req models.GetRangeOperationListReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetRangeOperationList(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetLastOperationListByAddr godoc
// @Summary last operation list by address
// @Schemes
// @Param JSON body models.GetLastOperationListByAddrReq true "GetLastOperationListByAddrReq params"
// @Description get last operation list page for address
// @Tags operation
// @Accept json
// @Produce json
// @Success 200 {object} models.GetLastOperationListByAddrResp
// @Failure 400
// @Router /operation/list/by-address/last [post]
func (r *Router) GetLastOperationListByAddr(c *gin.Context) {
	var req models.GetLastOperationListByAddrReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetLastOperationListByAddr(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetNextOperationListByAddr godoc
// @Summary next operation list by address
// @Schemes
// @Param JSON body models.GetNextOperationListByAddrReq true "GetNextOperationListByAddrReq params"
// @Description get next operation list page for address
// @Tags operation
// @Accept json
// @Produce json
// @Success 200 {object} models.GetNextOperationListByAddrResp
// @Failure 400
// @Router /operation/list/by-address/next [post]
func (r *Router) GetNextOperationListByAddr(c *gin.Context) {
	var req models.GetNextOperationListByAddrReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetNextOperationListByAddr(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetOperation godoc
// @Summary get operation
// @Schemes
// @Param JSON body models.GetOperationReq true "GetOperationReq params"
// @Description get operation full info
// @Tags operation
// @Accept json
// @Produce json
// @Success 200 {object} models.GetOperationResp
// @Failure 400
// @Router /operation/get [post]
func (r *Router) GetOperation(c *gin.Context) {
	var req models.GetOperationReq
	err := c.BindJSON(&req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := r.Core.GetOperation(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetPostgresqlRepoAlive
func (r *Router) GetPostgresqlRepoAlive(c *gin.Context) {
	resp, err := r.Core.GetPostgresqlRepoAlive()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetElasticsearchRepoAlive
func (r *Router) GetElasticsearchRepoAlive(c *gin.Context) {
	resp, err := r.Core.GetElasticsearchRepoAlive()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetServiceAlive
func (r *Router) GetServiceAlive(c *gin.Context) {
	resp, err := r.Core.GetServiceAlive()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
