package router

import (
	"healthcheck-monitor/internal/config"
	"healthcheck-monitor/internal/core"
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

	router.GET("/service/alive", r.GetServiceAlive)
	//router.StaticFile("/", "./../www/index.html")

	log.Fatal(router.Run(":" + r.Config.Monitor.Port))
}
