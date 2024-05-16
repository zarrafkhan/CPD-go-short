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
