package utils

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetIdHexInsertOneResult(result *mongo.InsertOneResult) (string, error) {
	if objId, ok := result.InsertedID.(primitive.ObjectID); ok {
		return objId.Hex(), nil
	}
	return "", fmt.Errorf("invalid insertedId %#v", result.InsertedID)
}
