package postgre

import (
	"github.com/azmash/go-clean/config"
	"github.com/azmash/go-clean/pkg/postgre"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

func NewDB(dbc config.DBConfig) *DB {
	db := postgre.ConnectDB(dbc)
	return &DB{
		db,
	}
}
