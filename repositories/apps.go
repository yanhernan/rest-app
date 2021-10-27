package repositories

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App struct {
	UID         primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Key         string             `bson:"key" json:"key"`
	Certificate string             `bson:"certificate" json:"certificate"`
	Created     time.Time          `bson:"created" json:"created"`
	Updated     time.Time          `bson:"updated" json:"updated"`
}

type _AppRepository struct {
	Repository
}

func Default() (*_AppRepository, error) {
	repository := &_AppRepository{}
	err := repository.setup()
	return repository, err
}
