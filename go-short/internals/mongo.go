// DB interaction

package internals

import (
	u "example/go-short/internals/util"
	"log"
	"os"

	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB(mongoKey string) *mongo.Client {
	//creates mongo client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoKey).SetServerAPIOptions(serverAPI)
	local, err := mongo.Connect(context.Background(), opts)
	u.Check(err)

	err = local.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	u.Check(err)

	fmt.Println("Mongo Bongo")
	return local
}

type URLList struct {
	Collection *mongo.Collection
}

// client global var - this func just returns the collection
func SetMongoColl(client *mongo.Client, name string, collName string) *mongo.Collection {
	collection := client.Database(name).Collection(collName)

	return collection
}

func AddURL(l *mongo.Collection, link string) interface{} {
	short := SetLink(link)
	res, e := l.InsertOne(context.Background(), short)
	u.Check(e)

	fmt.Println("Inserted a single document: ", res.InsertedID)
	return short
}

func GetLinkFromShort(l *mongo.Collection, short string) (string, error) {
	//convert mongo bson into struct type
	var url Link

	// filter := bson.D{{Key: "shortlink", Value: short}}
	filter := bson.M{"shortlink": short}

	e := l.FindOne(context.Background(), filter).Decode(&url)

	if e != nil {
		log.Println("Decode error: ", e)
	}

	//return full url
	return url.ID, nil
}

// returns DB close as error message
func Disc(client *mongo.Client) error {
	return client.Disconnect(context.Background())
}

func SetupMongo() (*mongo.Client, *mongo.Collection) {
	u.LoadEnv()
	mongoKey := os.Getenv("MONGO_KEY")
	name := os.Getenv("MONGO_NAME")
	collName := os.Getenv("MONGO_COLL_NAME")

	client := InitDB(mongoKey)
	coll := SetMongoColl(client, name, collName)

	fmt.Println("Mongo Collection Bongo")

	return client, coll
}
