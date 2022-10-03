package router

import (
	"log"
	"yolllo-manager/internal/config"
	"yolllo-manager/internal/core"

	docs "yolllo-manager/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	Config *config.Config
	Core   *core.Core
}

func NewRouter(cfg *config.Config, core *core.Core) (*Router, error) {

	return &Router{
		Config: cfg,
		Core:   core,
	}, nil
}

func (r *Router) Run() {
	router := gin.Default()

	// swagger
	if r.Config.Swagger.IsEnable {
		docs.SwaggerInfo.Title = "Yolllo-Manager API"
		docs.SwaggerInfo.Version = "0.1.0"
		docs.SwaggerInfo.Host = r.Config.Swagger.Host + ":" + r.Config.Router.Port
		docs.SwaggerInfo.BasePath = "/"
		docs.SwaggerInfo.Schemes = []string{"http", "https"}
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	// private api
	router.POST("/user/address/create", r.CreateUserAddress)
	router.POST("/user/transaction/create", r.CreateUserTransaction)
	router.POST("/user/staking/delegate", r.DelegateUserStaking)
	router.POST("/user/staking/reward/get", r.GetUserStakingReward)
	router.POST("/user/staking/reward/claim", r.ClaimUserStakingReward)
	router.POST("/user/staking/get", r.GetUserStaking)
	router.POST("/user/staking/undelegate", r.UndelegateUserStaking)
	router.POST("/user/staking/undelegated/get", r.GetUserStakingUndelegated)
	router.POST("/user/staking/undelegated/claim", r.ClaimUserStakingUndelegated)
	router.POST("/user/staking/total-stake/get", r.GetUserStakingTotalStake)
	router.POST("/user/staking/fee/get", r.GetUserStakingFee)
	router.POST("/user/staking/fee/set", r.SetUserStakingFee)
	// public api
	router.POST("/address/get", r.GetAddress)
	router.POST("/address/is-valid", r.IsValidAddress)
	router.POST("/transaction/get", r.GetTransaction)
	router.POST("/transaction/cost", r.GetTransactionCost)
	router.POST("/transaction/fee", r.GetTransactionFee)
	router.POST("/transaction/create", r.CreateTransaction)
	router.POST("/transaction/list/last", r.GetLastTransactionList)
	router.POST("/transaction/list/next", r.GetNextTransactionList)
	router.POST("/transaction/list/range", r.GetRangeTransactionList)
	router.POST("/transaction/list/by-address/last", r.GetLastTransactionListByAddr)
	router.POST("/transaction/list/by-address/next", r.GetNextTransactionListByAddr)
	router.POST("/block/by-nonce/get", r.GetBlockByNonce)
	router.POST("/block/by-hash/get", r.GetBlockByHash)
	router.POST("/block/last", r.GetLastBlock)
	router.POST("/block/list/last", r.GetLastBlockList)
	router.POST("/block/list/next", r.GetNextBlockList)
	router.POST("/staking/current-reward/monthly", r.GetStakingCurrentMonthlyReward)
	router.POST("/server/time/get", r.GetServerTime)

	log.Fatal(router.Run(":" + r.Config.Router.Port))
}
