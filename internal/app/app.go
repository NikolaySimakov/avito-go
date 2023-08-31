package app

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/NikolaySimakov/avito-go/config"
	"github.com/NikolaySimakov/avito-go/internal/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func databaseConnection() (*sql.DB, error) {
	// get connection info
	dbConfig := config.Database()

	// open DataBase
	db, err := sql.Open("postgres", dbConfig.GetURL())

	if err != nil {
		return nil, err
	}

	return db, nil
}

func runServer(router *mux.Router) {
	srv := &http.Server{
		Handler:      router,
		Addr:         config.Server().Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server will start at http://%s\n", config.Server().Address)
	log.Fatal(srv.ListenAndServe())
}

// Start HTTP server
func Run() {

	// Init DB
	database, err := databaseConnection()
	if err != nil {
		log.Fatal("Database error")
	}

	// Init repositories
	repos := db.NewRepositories(database)

	// Init router
	router := NewRouter(repos)

	// Init server
	runServer(router)
}
