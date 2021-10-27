package repositories

import (
	"github.com/yanhernan/rest-app/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	database *mongo.Database
}

var database *mongo.Database

func (repository *Repository) setup() error {
	if database != nil {
		repository.database = database
		return nil
	}
	db, err := utils.CreateDatabase()
	if err != nil {
		database = db
		repository.database = db
	}
	return err
}
