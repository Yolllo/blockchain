package main

import (
	"healthcheck-monitor/internal/config"
	"healthcheck-monitor/internal/core"
	"healthcheck-monitor/internal/router"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	core, err := core.NewCore(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	router, err := router.NewRouter(cfg, core)
	if err != nil {
		log.Fatal(err)
		return
	}

	router.Run()
}
