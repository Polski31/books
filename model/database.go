package model

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collection *mongo.Collection
}

var MngInst MongoInstance

func ConnectToDatabase() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(databaseUri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)

	if err != nil {
		fmt.Printf(err.Error(), log.Default())
	}

	MngInst.Client = client
	MngInst.DB = client.Database(databaseName)
	MngInst.Collection = MngInst.DB.Collection(databaseCollectionName)
}
