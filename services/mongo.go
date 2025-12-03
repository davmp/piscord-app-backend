package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoService struct {
	client   *mongo.Client
	database *mongo.Database
	uri      string
}

func NewMongoService(uri string) *MongoService {
	return &MongoService{
		uri: uri,
	}
}

func (m *MongoService) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(m.uri)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	m.client = client
	m.database = client.Database("piscord")

	// m.createIndexes()

	return nil
}

func (m *MongoService) Disconnect() error {
	if m.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return m.client.Disconnect(ctx)
	}
	return nil
}

func (m *MongoService) GetDatabase() *mongo.Database {
	return m.database
}

func (m *MongoService) GetCollection(name string) *mongo.Collection {
	return m.database.Collection(name)
}
