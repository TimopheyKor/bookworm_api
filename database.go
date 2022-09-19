package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Define a struct to hold the database client.
type mongoDB struct {
	client *mongo.Client
}

// Define constant variables for accessing various parts of the database.
const (
	dbLocalData        = "local"
	dbAdmins           = "admin"
	dbConfiguration    = "config"
	activeCollection   = "active_books"
	wishlistConnection = "wishlist"
	logsCollection     = "startup_log"
)

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

// addBook takes a bookData struct and a context, then attempts to add it as a
// document to the active_books collection of the database.
func (db *mongoDB) addBook(document bookData, ctx context.Context) error {
	coll := db.client.Database(dbLocalData).Collection(activeCollection)
	result, err := coll.InsertOne(ctx, document)
	id := result.InsertedID
	fmt.Println("Inesrtion ID: ", id)
	return err
}

// getAllBooks takes a request context, then creates a cursor with which it
// iterates over all available documents in the active_books collection
// of the database, populates a slice with them, and returns the slice.
func (db *mongoDB) getAllBooks(ctx context.Context) ([]bson.D, error) {
	coll := db.client.Database(dbLocalData).Collection(activeCollection)

	// Attempt to aggregate the data into a cursor variable.
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	// Populate a slice with all the book data using the cursor.
	var allBooks []bson.D
	err = cursor.All(ctx, &allBooks)
	if err != nil {
		return nil, err
	}
	return allBooks, nil
}

func getBook() {

}

func updateBook() {

}
