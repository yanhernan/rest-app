package utils

import (
	"context"
	"time"

	"github.com/yanhernan/rest-app/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDatabase() (*mongo.Database, error) {
	appConfig := config.GetAppConfig()
	dataStore := appConfig.DataStorage
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	options := options.Client().ApplyURI(dataStore.URI)
	options.SetMaxPoolSize(10)
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, err
	}

	return client.Database(dataStore.DataBaseName), nil
}
