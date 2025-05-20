package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type StorageMongoImpl struct {
	mongoClient *mongo.Client
	collection  *mongo.Collection
}

func NewStorage(input *mongo.Client) Storage {
	return &StorageMongoImpl{
		mongoClient: input,

	}
}

func NewStorageMongoImpl(db *mongo.Database, collectionName string) *StorageMongoImpl {
	
	return &StorageMongoImpl{
		collection: db.Collection(collectionName),
	}

}
