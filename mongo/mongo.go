package mongo

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type M bson.M

func NewMongoClient(ctx context.Context, uri string) *mongo.Client {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client
}

func Disconnect(ctx context.Context, client *mongo.Client) error {
	log.Printf("Disconnecting from MongoDB")
	return client.Disconnect(ctx)
}

func Insert(col *mongo.Collection, data interface{}) error {
	_, err := col.InsertOne(context.TODO(), data)
	return err
}

// out - a pointer to the type you want to decode into
func Upsert(col *mongo.Collection, filter M, data interface{}, out interface{}) error {
	res, err := col.ReplaceOne(context.TODO(), filter, data, options.Replace().SetUpsert(true))
	if err != nil {
		return err
	}
	if res.UpsertedID != nil {
		return FindOne(col, M{"_id": res.UpsertedID}, out)
	}
	return nil
}

// out - a pointer to the type you want to decode into
func Update(col *mongo.Collection, filter M, data interface{}, out interface{}) error {
	res, err := col.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		return err
	}
	if res.UpsertedID != nil {
		return FindOne(col, M{"_id": res.UpsertedID}, out)
	}
	return nil
}

// out - a pointer to the type you want to decode into incase you want to return the deleted document
func Delete(col *mongo.Collection, filter M, out interface{}) error {
	err := FindOne(col, filter, out)
	if err != nil {
		return errors.New("delete error: no document found")
	}
	_, err = col.DeleteOne(context.TODO(), filter)
	return err
}

// out - a pointer to the type you want to decode into
func FindOne(col *mongo.Collection, filter M, out interface{}) error {
	return col.FindOne(context.TODO(), filter).Decode(out)
}

// out - a pointer to a slice of the type you want to decode into
func Find(col *mongo.Collection, filter M, out interface{}) error {
	if cur, err := col.Find(context.TODO(), filter); err == nil {
		return cur.All(context.TODO(), out)
	} else {
		return err
	}
}

// out - a pointer to a slice of the type you want to decode into
func Aggregate(col *mongo.Collection, pipe []M, out interface{}) error {
	if cur, err := col.Aggregate(context.TODO(), pipe); err == nil {
		return cur.All(context.TODO(), out)
	} else {
		return err
	}
}
