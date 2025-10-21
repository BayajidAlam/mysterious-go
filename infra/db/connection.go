package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.mod/config"
)

func GetConnectionString(cnf *config.DBConfig) string {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		cnf.HOST,
		cnf.PORT,
		cnf.USER,
		cnf.PASSWORD,
		cnf.NAME,
	)
	if !cnf.EnableSSLMODE {
		connString += " sslmode=disable"
	}
	return connString
}

func NewConnection(cnf *config.DBConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(cnf)
	dbCon, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
