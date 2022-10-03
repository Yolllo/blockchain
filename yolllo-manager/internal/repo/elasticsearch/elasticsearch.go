package elasticsearch

import (
	"yolllo-manager/internal/config"

	es "github.com/elastic/go-elasticsearch/v8"
)

type Repository struct {
	Config *config.Config
	Conn   *es.Client
}
