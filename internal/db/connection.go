package db

import (
	"database/sql"
	"log"

	"github.com/NikolaySimakov/avito-go/config"
	_ "github.com/lib/pq"
)

func Database() *sql.DB {
	// get connection info
	dbConfig := config.Database()

	// open DataBase
	db, err := sql.Open("postgres", dbConfig.GetURL())

	if err != nil {
		log.Fatal(err)
	}

	return db
}
