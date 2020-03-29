package database

import (
	"context"

	utils "golang-book-api/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB represents database object
var DB *mongo.Database

// ConnectToDB establish connection to MongoDB
func ConnectToDB(connectionString string, dbName string) {
	// MongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionString))
	if err != nil {
		utils.Logger.Error(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		utils.Logger.Error(err)
	}
	DB = client.Database(dbName)
	utils.Logger.Debugf("Connceted do database %s", DB.Name())
}

// GetCollection retrives collection with given name from database
func GetCollection(collectionName string) *mongo.Collection {
	collection := DB.Collection(collectionName)
	return collection
}
