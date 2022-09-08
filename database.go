package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	client *mongo.Client
}

// newMongoDB takes a context and a URI, then attempts to initialize a new
// MongoDB client. It returns a new mongoDB object holding the client, and
// an error.
func newMongoDB(ctx context.Context, uri string) (*mongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return &mongoDB{client: client}, nil
}

// disconnect takes a context, then attempts to disconnect the database's
// client. It returns an error if it fails.
func (db *mongoDB) disconnect(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		return err
	}
	return nil
}

func addBook() {

}

func getBook() {

}

func updateBook() {

}
