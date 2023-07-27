package database

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/yuttasakcom/go-hexa/app/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Database
}

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

func (m *MongoDB) Create(v interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	structName := strings.ToLower(reflect.TypeOf(v).Elem().Name())
	pluralize := pluralize.NewClient()
	collectionName := pluralize.Plural(structName)
	_, err := m.Collection(collectionName).InsertOne(ctx, v)
	if err != nil {
		return err
	}
	return nil
}
