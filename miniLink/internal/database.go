package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoConnection *mongo.Database
)

func BootstrapMongo(uri, dbName string, timeout time.Duration) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		return nil, err
	}

	mongoConnection = client.Database(dbName)

	return client.Database(dbName), nil
}

func GetMongo() *mongo.Database {
	return mongoConnection
}

func CloseMongo(db *mongo.Database) error {
	return db.Client().Disconnect(context.Background())
}
