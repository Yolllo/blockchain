package repo

import (
	"context"
	"yolllo-manager/internal/config"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	Config *config.Config
	Conn   *pgx.Conn
}

func NewRepository(cfg *config.Config) (repo *Repository, err error) {
	repo = &Repository{}
	repo.Config = cfg

	databaseUrl := "postgres://" + cfg.Repo.User + ":" + cfg.Repo.Password + "@" + cfg.Repo.Host + ":" + cfg.Repo.Port + "/" + cfg.Repo.Name
	repo.Conn, err = pgx.Connect(context.Background(), databaseUrl)
	if err != nil {

		return
	}

	return
}
