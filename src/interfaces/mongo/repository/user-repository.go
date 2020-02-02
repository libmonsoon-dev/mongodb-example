package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"mongodb-example/src/domain"
	"mongodb-example/src/interfaces/mongo/utils"
)

const UsersCollectionName = "users"

type MongoUserRepository struct {
	db *mongo.Database
}

func NewMongoUserRepository(db *mongo.Database) *MongoUserRepository {
	return &MongoUserRepository{db}
}

func (m MongoUserRepository) getCollection() *mongo.Collection {
	return m.db.Collection(UsersCollectionName)
}

func (m MongoUserRepository) Store(ctx context.Context, user domain.User) (string, error) {
	result, err := m.getCollection().InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	return utils.GetIdHexInsertOneResult(result)
}
