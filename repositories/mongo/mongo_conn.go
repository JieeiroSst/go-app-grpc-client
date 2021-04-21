package mongo 

import (
	"context"
	"fmt"
	"sync"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mutex    sync.Mutex
	instance *Mongoconn
)

type Mongoconn struct {
	collection *mongo.Collection
}

type Config struct {
	DSN        string
	DB         string
	Collection string
}

func GetMongoConnInstance(cf Config) *Mongoconn {
	fmt.Println(cf)
	if instance == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance == nil {
			dsn := cf.DSN
			clientOptions := options.Client().ApplyURI(dsn)
			client, err := mongo.Connect(context.TODO(), clientOptions)
			if err != nil {
				panic(err)
			}

			if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
				panic(err)
			}
			instance = &Mongoconn{
				collection: client.Database(cf.DB).Collection(cf.Collection),
			}
		}
	}
	return instance
}

func NewMongoSqlRepo(cf Config) *Mongoconn {
	return &Mongoconn{
		collection: GetMongoConnInstance(cf).collection,
	}
}