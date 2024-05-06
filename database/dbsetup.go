package database

import (
	"context"
	"log"
	"os"
	"restraument-management/custom"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBSetup()

func DBSetup() *mongo.Client {

	godotenv.Load()

	clientoption := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.Background(), clientoption)

	if err != nil {
		log.Panicln(custom.ErrDBNotConnected)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Panicln(custom.ErrDBNotPinnged)
	}
	log.Println(custom.DBConnected)
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	godotenv.Load()
	var collection *mongo.Collection = client.Database(os.Getenv("DB_NAME")).Collection(collectionName)
	return collection
}
