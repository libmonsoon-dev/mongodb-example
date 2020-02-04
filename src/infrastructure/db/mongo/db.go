package mongo

import "go.mongodb.org/mongo-driver/mongo"

func NewDbConnection(config *Config) (*mongo.Database, error) {
	client, err := NewClient(config)

	if err != nil {
		return nil, err
	}

	return client.Database(config.DataBaseName), nil
}

func MustNewDbConnection(config *Config) *mongo.Database {
	db, err := NewDbConnection(config)
	if err != nil {
		panic(err)
	}
	return db
}
