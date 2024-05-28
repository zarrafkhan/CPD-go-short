// DB interaction

// TESTED NEW TOPOLOGY

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

var mctx, _ = context.WithTimeout(context.Background(), 30*time.Second)

func InitDB(mongoKey string) *mongo.Client {
	//creates mongo client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(mongoKey).SetServerAPIOptions(serverAPI)
	local, err := mongo.Connect(context.Background(), opts)
	u.Check(err)

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

func AddURL(l *mongo.Collection, link string, ctx context.Context) (string, string) {

	short := SetLink(link)
	_, e := l.InsertOne(ctx, short)
	u.Check(e)

	return short.ID, short.ShortLink
}

// func AddMulti(l *mongo.Collection, links []string) {
// 	if len(links) > 0 {
// 		for i := 0; i < len(links); i++ {
// 			AddURL(l, links[i], con)
// 		}
// 	}
// }

func DeletURL(l *mongo.Collection, link string) error {
	filter := bson.M{"id": link}
	return l.FindOneAndDelete(mctx, filter).Err()
}

func GetLinkFromShort(l *mongo.Collection, short string) (string, error) {
	//convert mongo bson into link struct type
	var url Link
	filter := bson.M{"shortlink": short}

	e := l.FindOne(context.TODO(), filter).Decode(&url)
	if e != nil {
		log.Println("Decode error: ", e)
	}

	return url.ID, nil
}

func CountDocs(l *mongo.Collection) (int64, error) {
	opts := options.Count()
	shorts, e := l.CountDocuments(context.TODO(), bson.D{}, opts)

	if e != nil {
		u.Check(e)
	}
	return shorts, e
}

func GetShorts(l *mongo.Collection) []Link {
	var list []Link
	// Define the filter to find all documents
	findOptions := options.Find()

	cur, err := l.Find(mctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Link
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())

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

	//ctx, cancel := context.WithTimeout(context.Background(),10* time.Second)
	return client, coll
}
