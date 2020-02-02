package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"mongodb-example/src/domain"
	"mongodb-example/src/interfaces/mongo/utils"
)

const GamesCollectionName = "games"

type MongoGameRepository struct {
	db *mongo.Database
}

func NewMongoGameRepository(db *mongo.Database) *MongoGameRepository {
	return &MongoGameRepository{db}
}

func (m MongoGameRepository) getCollection() *mongo.Collection {
	return m.db.Collection(GamesCollectionName)
}

func (m MongoGameRepository) Store(ctx context.Context, user domain.Game) (string, error) {
	result, err := m.getCollection().InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return utils.GetIdHexInsertOneResult(result)
}
