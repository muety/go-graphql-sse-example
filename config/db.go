package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

const (
	CollectionOrders   = "orders"
	CollectionProducts = "products"
)

var (
	client      *mongo.Client
	db          *mongo.Database
	collections map[string]*mongo.Collection
)

func InitDbSync() (*mongo.Database, error) {
	log.Println("connecting to database")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Db.Url)); err == nil {
		client = mongoClient
	} else {
		return nil, err
	}

	// retry 10 times
	for i := 0; i < 10; i++ {
		if err := client.Ping(ctx, readpref.Primary()); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	log.Println("successfully connected to database")

	db = client.Database(config.Db.Name)
	collections = make(map[string]*mongo.Collection)

	return db, nil
}

func GetDb() *mongo.Database {
	return db
}

func GetDbCollection(key string) *mongo.Collection {
	if _, ok := collections[key]; !ok {
		collections[key] = db.Collection(key)
	}
	return collections[key]
}
