package crud

import (
	"log"

	"github.com/NikolaySimakov/avito-go/internal/db"
	"github.com/NikolaySimakov/avito-go/internal/models"
)

var database = db.Database()

func CreateUser(user models.User) {
	_, err := database.Exec("INSERT INTO users (id, segments) VALUES ($1, $2)", user.Id, user.Segments.ToString())
	if err != nil {
		log.Fatal(err)
	}
}
