// DB interaction

package libraries

import (
	u "example/go-short/libraries/util"
	"log"
	"os"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)

func InitDB(mongoKey string) *mongo.Client {
	//creates mongo client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoKey).SetServerAPIOptions(serverAPI)
	local, err := mongo.Connect(context.Background(), opts)
	u.Check(err)

	//err = local.Database("admin").RunCommand(mctx, bson.D{{Key: "ping", Value: 1}}).Err()
	err = local.Ping(context.TODO(), nil)
	u.Check(err)

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

func AddURL(l *mongo.Collection, link string) (string, string) {
	short := SetLink(link)
	_, e := l.InsertOne(context.Background(), short)
	u.Check(e)
	return short.ID, short.ShortLink
}

func DeletURL(l *mongo.Collection, link string) error {
	filter := bson.M{"id": link}
	return l.FindOneAndDelete(mctx, filter).Err()
}

func GetLinkFromShort(l *mongo.Collection, short string) (string, error) {
	//convert mongo bson into struct type
	var url Link

	// filter := bson.D{{Key: "shortlink", Value: short}}
	filter := bson.M{"shortlink": short}

	e := l.FindOne(mctx, filter).Decode(&url)

	if e != nil {
		log.Println("Decode error: ", e)
	}

	//return full url
	return url.ID, nil
}

// returns DB close as error message
func Disc(client *mongo.Client) error {
	return client.Disconnect(mctx)
}

func SetupMongo() (*mongo.Client, *mongo.Collection) {
	u.LoadEnv()
	mongoKey := os.Getenv("MONGO_KEY")
	name := os.Getenv("MONGO_NAME")
	collName := os.Getenv("MONGO_COLL_NAME")
	client := InitDB(mongoKey)
	coll := SetMongoColl(client, name, collName)

	return client, coll
}
