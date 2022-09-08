package main

import (
	"context"
	"fmt"

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

// TODO: Change this to use a passed-down gin context.
// addBook takes a bookData struct and a context, then attempts to add it as a
// document to the active_books collection of the database.
func (db *mongoDB) addBook(document bookData, ctx context.Context) error {
	coll := db.client.Database("local").Collection("active_books")
	result, err := coll.InsertOne(ctx, document)
	id := result.InsertedID
	fmt.Println("Inesrtion ID: ", id)
	return err
}

func getBook() {

}

func updateBook() {

}
