package repo

import (
	"context"
	"fmt"
	"log"
	"yolllo-manager/internal/config"
	"yolllo-manager/internal/repo/elasticsearch"
	"yolllo-manager/internal/repo/postgresql"

	es "github.com/elastic/go-elasticsearch/v8"

	"github.com/jackc/pgx/v4"
)

type RepositoryManager struct {
	Config *config.Config
	PG     *postgresql.Repository
	ES     *elasticsearch.Repository
}

func NewRepositoryManager(cfg *config.Config) (repo *RepositoryManager, err error) {
	repo = &RepositoryManager{
		Config: cfg,
	}

	// init PostgresSQL
	log.Println("Init PostgresSQL repository")
	repo.PG = &postgresql.Repository{
		Config: cfg,
	}
	pgFullUrl := "postgres://" + cfg.Repo.PostgreSQL.User + ":" + cfg.PostgreSQLPassword + "@" + cfg.Repo.PostgreSQL.Host + ":" + cfg.Repo.PostgreSQL.Port + "/" + cfg.Repo.PostgreSQL.Name
	repo.PG.Conn, err = pgx.Connect(context.Background(), pgFullUrl)
	if err != nil {

		return
	}

	// init ElasticSearch
	log.Println("Init ElasticSearch repository")
	repo.ES = &elasticsearch.Repository{
		Config: cfg,
	}
	esFullUrl := "http://" + cfg.Repo.ElasticSearch.Host + ":" + cfg.Repo.ElasticSearch.Port
	repo.ES.Conn, err = es.NewClient(es.Config{
		Addresses: []string{
			esFullUrl,
		},
		Username: cfg.Repo.ElasticSearch.User,
		Password: cfg.ElasticSearchPassword,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
