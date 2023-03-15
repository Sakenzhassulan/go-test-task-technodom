package db

import (
	"context"
	"github.com/Sakenzhassulan/go-test-task-technodom/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Instance struct {
	Collection *mongo.Collection
}

func NewDbCollection(config *config.Config) (*Instance, error) {
	clientOptions := options.Client().ApplyURI(config.DbUri)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database(config.DbName).Collection(config.DbCollection)
	return &Instance{Collection: collection}, nil
}
