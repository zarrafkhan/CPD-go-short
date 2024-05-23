// DB interaction

package libraries

import (
	u "example/go-short/libraries/util"
	utils "example/go-short/libraries/util"
	"log"
	"os"
	"time"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mctx, _ = context.WithTimeout(context.Background(), 30*time.Second)

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

func AddMulti(l *mongo.Collection, links []string) {
	if len(links) > 0 {
		for i := 0; i < len(links); i++ {
			AddURL(l, links[i])
		}
	}
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

func CountDocs(l *mongo.Collection) (int64, error) {
	opts := options.Count()
	shorts, e := l.CountDocuments(context.TODO(), bson.D{}, opts)

	if e != nil {
		utils.Check(e)
	}
	return shorts, e
}

func GetShorts(l *mongo.Collection) []string {
	var list []string
	// Define the filter to find all documents
	filter := bson.M{}

	// Define the projection to only include the "ShortLink" field
	projection := bson.M{"ShortLink": 1, "_id": 0}

	// Find all documents with the specified filter and projection
	cursor, err := l.Find(context.TODO(), filter, options.Find().SetProjection(projection))
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	// Iterate through the cursor and extract the "ShortLink" field values
	for cursor.Next(context.TODO()) {
		var result bson.M
		if err = cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}

		if shortLink, ok := result["ShortLink"].(string); ok {
			list = append(list, shortLink)
		}
	}

	if err = cursor.Err(); err != nil {
		log.Fatal(err)
	}

	return list
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
