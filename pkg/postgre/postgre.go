package postgre

import (
	"fmt"
	"log"

	"github.com/azmash/go-clean/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDB(dbc config.DBConfig) *sqlx.DB {
	postgreInfo := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		dbc.Host,
		dbc.Port,
		"postgres",
		dbc.Credential,
		"postgres",
	)

	sqlCon, err := sqlx.Open("postgres", postgreInfo)
	if err != nil {
		log.Println("Err Init DB: ", err)
	}

	return sqlCon
}
