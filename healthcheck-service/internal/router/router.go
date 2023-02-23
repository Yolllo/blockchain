package router

import (
	"healthcheck-service/internal/config"
	"healthcheck-service/internal/core"
	"log"

	"github.com/gin-gonic/gin"
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

	router.GET("/node/status", r.GetNodeStatus)
	//router.POST("/node/exec", r.ExecNode)
	router.GET("/service/alive", r.GetServiceAlive)
	log.Fatal(router.Run(":" + r.Config.ServicePort))
}
