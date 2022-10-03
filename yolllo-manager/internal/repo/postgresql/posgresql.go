package postgresql

import (
	"yolllo-manager/internal/config"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	Config *config.Config
	Conn   *pgx.Conn
}
