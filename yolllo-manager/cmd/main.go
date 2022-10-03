package main

import (
	"log"
	"yolllo-manager/internal/config"
	"yolllo-manager/internal/core"
	"yolllo-manager/internal/repo"
	"yolllo-manager/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = cfg.EnterInitData()
	if err != nil {
		log.Fatal(err)
		return
	}

	repo, err := repo.NewRepositoryManager(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}

	core, err := core.NewCore(cfg, repo)
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
