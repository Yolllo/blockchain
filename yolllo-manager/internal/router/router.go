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
	// public api
	router.POST("/address/get", r.GetAddress)
	router.POST("/transaction/get", r.GetTransaction)
	router.POST("/transaction/cost", r.GetTransactionCost)
	router.POST("/transaction/create", r.CreateTransaction)
	router.POST("/block/by-nonce/get", r.GetBlockByNonce)
	router.POST("/block/by-hash/get", r.GetBlockByHash)
	router.POST("/block/last", r.GetLastBlock)

	log.Fatal(router.Run(":" + r.Config.Router.Port))
}
