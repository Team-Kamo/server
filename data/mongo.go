package data

import (
	"OctaneServer/config"
	"context"
	"fmt"

	"github.com/goccy/go-json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Mongo struct {
	Type   string
	Ready  bool
	client *mongo.Client
}

func (db *Mongo) Open() error {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.CurrentConfig.Db.Uri))
	if err != nil {
		return err
	}
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	db.client = client
	db.Ready = true
	db.Type = MongoDB
	return nil
}

func (db *Mongo) Close() error {
	return db.client.Disconnect(context.TODO())
}

func (db *Mongo) Insert(table string, key string, data interface{}) error {
	_, err := db.client.Database(config.CurrentConfig.Db.Name).Collection(table).InsertOne(context.TODO(), bson.D{{Key: "key", Value: key}, {Key: "value", Value: data}})
	if err != nil {
		return err
	}
	return nil
}

func (db *Mongo) Update(table string, key string, data interface{}) error {
	fmt.Print(data)
	_, err := db.client.Database(config.CurrentConfig.Db.Name).Collection(table).ReplaceOne(context.TODO(), bson.D{{Key: "key", Value: key}}, bson.D{{Key: "key", Value: key}, {Key: "value", Value: data}})
	if err == mongo.ErrNoDocuments {
		return ErrNoEntry
	} else if err != nil {
		return err
	}
	return nil
}

func (db *Mongo) Get(table string, key string) ([]byte, error) {
	var result bson.M
	res := db.client.Database(config.CurrentConfig.Db.Name).Collection(table).FindOne(context.TODO(), bson.D{{Key: "key", Value: key}})
	err := res.Decode(&result)
	if err == mongo.ErrNoDocuments {
		return nil, ErrNoEntry
	} else if err != nil {
		return nil, err
	}
	j, err := json.Marshal(result["value"])
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (db *Mongo) Delete(table string, key string) error {
	result, err := db.client.Database(config.CurrentConfig.Db.Name).Collection(table).DeleteOne(context.TODO(), bson.D{{Key: "key", Value: key}})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return ErrNoEntry
	}
	return nil
}

func (db *Mongo) Drop(table string) error {
	err := db.client.Database(config.CurrentConfig.Db.Name).Collection(table).Drop(context.TODO())
	if err != nil {
		return err
	}
	return nil
}
