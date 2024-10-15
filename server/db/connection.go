package db

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func ConnectDB() *mongo.Database {
	err := godotenv.Load(".env")
	if os.Getenv("ENV") != "production" {
		if err != nil {
			log.Fatal("Error loading environment file: ", err)
		}
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOption := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal("Error from MongoDB Connection")
	}
	err = client.Ping(context.Background(), nil)
	database := client.Database("golang_entire_db")
	Db = database
	return Db
}
