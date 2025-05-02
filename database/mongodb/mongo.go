package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MOngoConn() (*mongo.Client, error) {
	url := os.Getenv("mongoURL")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Panic(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		return nil, fmt.Errorf("Database Ping failrd: %v", err)
	}
	fmt.Println("Connected to MongoDB with database/MongoDB")
	return client, err
}
