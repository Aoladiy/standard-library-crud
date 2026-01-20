package db

import (
	"database/sql"
	"log"

	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitDB(v envLoading.EnvVariables) *sql.DB {
	var err error
	DB, err = sql.Open("pgx", v.GetDsn())
	if err != nil {
		log.Fatalln("Cannot connect to database:", err)
		return nil
	}
	if err = DB.Ping(); err != nil {
		log.Fatalln("Cannot ping database:", err)
		return nil
	}
	return DB
}
