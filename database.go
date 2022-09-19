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
func (db *mongoDB) getAllBooks(ctx context.Context) ([]bookData, error) {
	coll := db.client.Database(dbLocalData).Collection(activeCollection)

	// Attempt to aggregate the data into a cursor variable.
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Populate a slice with all the book data using the cursor. cursor.All()
	// automatically unmarshalls the bson data into bookData.
	var allBooks []bookData
	err = cursor.All(ctx, &allBooks)
	if err != nil {
		return nil, err
	}
	return allBooks, nil
}

// getBook takes the name of a book as a string and the request context, then
// attempts to find a book with that title in the databse, returning the book's
// data if it is successful.
func (db *mongoDB) getBook(name string, ctx context.Context) (*bookData, error) {
	coll := db.client.Database(dbLocalData).Collection(activeCollection)

	var book bookData
	err := coll.FindOne(ctx, bson.D{{"title", name}}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrBookNotFound
		}
		return nil, err
	}
	return &book, nil
}

func updateBook() {

}
