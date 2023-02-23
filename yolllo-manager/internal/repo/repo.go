package repo

import (
	"context"
	"fmt"
	"log"
	"time"
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

	go repo.PingPG(pgFullUrl)
	go repo.PingES(esFullUrl, cfg)

	return
}

func (repo *RepositoryManager) PingPG(connUrl string) {
	for {
		time.Sleep(3 * time.Second)

		ctx := context.Background()

		err := repo.PG.Conn.Ping(ctx)
		if err != nil {
			fmt.Println("PG RECCONECTION...")
			conn, err := pgx.Connect(context.Background(), connUrl)
			if err == nil {
				repo.PG.Conn = conn
				fmt.Println("PG RECONNECTED!")
			}
		}
	}
}

func (repo *RepositoryManager) PingES(connUrl string, cfg *config.Config) {
	for {
		time.Sleep(3 * time.Second)

		ctx := context.Background()

		pingRequest := repo.ES.Conn.Ping.WithContext(ctx)
		_, err := repo.ES.Conn.Ping(pingRequest)
		if err != nil {
			fmt.Println("ES RECCONECTION...")
			conn, err := es.NewClient(es.Config{
				Addresses: []string{
					connUrl,
				},
				Username: cfg.Repo.ElasticSearch.User,
				Password: cfg.ElasticSearchPassword,
			})
			if err == nil {
				repo.ES.Conn = conn
			}
		}
	}
}
