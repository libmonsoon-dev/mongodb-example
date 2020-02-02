package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewClient(config *Config) (*mongo.Client, error) {
	uri := fmt.Sprintf("mongodb://%v:%v", config.Host, config.Port)
	opts := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		return nil, fmt.Errorf("mongo connect error: %w", err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("mongo ping error: %w", err)
	}

	return client, nil
}
