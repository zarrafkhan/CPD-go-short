// DB interaction

package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// handle error
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Init(mongoKey string, db string) error {
	//creates mongo client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoKey).SetServerAPIOptions(serverAPI)
	local, err := mongo.Connect(context.Background(), opts)
	Check(err)

	client = local

	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	Check(err)

	return err
}

func Disc() error {
	return client.Disconnect(context.Background())
}
