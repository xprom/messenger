package repository

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	logger         *log.Logger
	client         *mongo.Client
	testCollection *mongo.Collection
}

func New(client *mongo.Client, logger *log.Logger) *Repository {
	createIndexOption := options.Index()
	createIndexOption.SetUnique(true)

	repo := &Repository{
		logger:         logger,
		client:         client,
		testCollection: client.Database("messenger").Collection("messenger_test"),
	}
	return repo
}
