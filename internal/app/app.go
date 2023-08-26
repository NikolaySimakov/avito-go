package app

import (
	"log"
	"net/http"
	"time"

	"github.com/NikolaySimakov/avito-go/config"
)

// Start HTTP server
func Run() {
	// Init router
	router := InitRouter()

	// Init server
	srv := &http.Server{
		Handler:      router,
		Addr:         config.Server().Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
