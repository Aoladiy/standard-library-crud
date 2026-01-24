package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/Aoladiy/standard-library-crud/internal/envLoading"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitDB(v envLoading.EnvVariables) *sql.DB {
	DB, err := sql.Open("pgx", v.GetDsn())
	if err != nil {
		log.Fatalln("Cannot connect to database:", err)
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	if err = DB.PingContext(ctx); err != nil {
		log.Fatalln("Cannot ping database:", err)
		return nil
	}
	return DB
}
