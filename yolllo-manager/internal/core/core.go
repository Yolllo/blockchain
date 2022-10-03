package core

import (
	"math/rand"
	"time"
	"yolllo-manager/internal/config"
	"yolllo-manager/internal/repo"
)

type Core struct {
	Config   *config.Config
	Repo     *repo.RepositoryManager
	Mnemonic string
}

func NewCore(cfg *config.Config, repo *repo.RepositoryManager) (core *Core, err error) {
	core = &Core{
		Config: cfg,
		Repo:   repo,
	}
	rand.Seed(time.Now().UnixNano())

	return
}
