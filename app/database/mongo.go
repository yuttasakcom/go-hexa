package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yuttasakcom/go-hexa/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(config config.IDb) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Url()))
	if err != nil {
		log.Fatalf("Error connect mongo %v", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client.Database(config.Dbname())
}
