// DB interaction

package internals

import (
	u "example/go-short/internals/util"
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
func GetMongoColl(client *mongo.Client, name string, collName string) *mongo.Collection {
	collection := client.Database(name).Collection(collName)

	return collection
}

func (l *URLList) AddURL(link string) (interface{}, error) {
	short := SetLink(link)
	res, e := l.Collection.InsertOne(context.Background(), short)
	u.Check(e)
	return res, nil
}

func InsertURL(l *mongo.Collection, url string) string {
	result := SetLink(url)
	insertResult, err := l.InsertOne(context.TODO(), result)
	u.Check(err)

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	return result.ShortLink
}

func FindShortByLink(l *mongo.Collection, url_id string) (string, error) {
	//convert mongo bson into struct type
	var url Link

	e := l.FindOne(context.Background(),
		bson.D{{Key: "id", Value: "url_id"}}).Decode(&url)
	u.Check(e)

	return url.ShortLink, nil
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
	coll := GetMongoColl(client, name, collName)

	fmt.Println("Mongo Collection Bongo")

	return client, coll
}
