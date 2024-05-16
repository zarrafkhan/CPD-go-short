// DB interaction

package internals

import (
	"context"
	u "example/go-short/internals/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type URLList struct {
	Collection *mongo.Collection
}

func (l *URLList) AddURL(link *Link) (interface{}, error) {
	res, e := l.Collection.InsertOne(context.Background(), link)
	u.Check(e)
	return res, nil
}

func (l *URLList) FindLinkByID(url_id string) (*Link, error) {
	//convert bson into struct type
	var url Link

	e := l.Collection.FindOne(context.Background(),
		bson.D{{Key: "id", Value: "url_id"}}).Decode(&url)
	u.Check(e)

	return &url, nil
}

// func (l *URLList) ReplaceLinksByID(url_id string, link *models.Url) (int32, error){
// 	var links int32
// 	res, e := l.Collection.InsertOne(context.Background(), link)
// 	u.Check(e)
// 	return links, nil

// }

// var client *mongo.Client

// func Init(mongoKey string, dbs string) error {
// 	//creates mongo client
// 	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
// 	opts := options.Client().ApplyURI(mongoKey).SetServerAPIOptions(serverAPI)
// 	local, err := mongo.Connect(context.Background(), opts)
// 	u.Check(err)

// 	client = local

// 	err = client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
// 	u.Check(err)

// 	return err
// }

// func Disc() error {
// 	return client.Disconnect(context.Background())
// }
