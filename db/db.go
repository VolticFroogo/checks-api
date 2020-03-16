package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri = os.Getenv("DB")

	Customer, Server, Log, Operator *mongo.Collection
)

// Init initialises the database.
func Init() (err error) {
	opts := options.Client().ApplyURI(uri)

	client, err := mongo.NewClient(opts)
	if err != nil {
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return
	}

	err = client.Ping(ctx, opts.ReadPreference)
	if err != nil {
		return
	}

	db := client.Database("checks")

	Customer = db.Collection("customer")
	Server = db.Collection("server")
	Log = db.Collection("log")
	Operator = db.Collection("operator")
	return
}
