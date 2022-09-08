package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Router to access a controller and a logger if necessary. The router should
// handle all http & gin methods.
type router struct {
	c controller
}

// getBook takes the name of a book from the request context, and returns
// a JSON response or an error based on if the book exists in the dataset or not.
func (r *router) getBook(c *gin.Context) {
	name := c.Param("name")
	book, err := r.c.getBook(name)
	if err != nil {
		c.Status(interpretError(err))
		return
	}
	c.JSONP(http.StatusOK, book)
}

// getAllBooks returns a JSON response containing a map of all the books
// currently inside the database.
func (r *router) getAllBooks(c *gin.Context) {
	c.JSONP(http.StatusOK, r.c.getAllBooks())
}

// addBook adds an empty book except for the provided title to the database.
func (r *router) addBook(c *gin.Context) {
	name := c.Param("name")
	err := r.c.addNewBook(name)
	if err != nil {
		c.JSONP(interpretError(err), fmt.Sprintf("Error: %v", err))
		return
	}
	c.Status(http.StatusCreated)
}

// addBookWithBody checks to see if a JSON body has been given to the request,
// then checks if the body is valid. If so, it calls addBook to make a new
// entry into the database, then populates the book with the JSON body data.
func (r *router) addBookWithBody(c *gin.Context) {
	var json bookData
	name := c.Param("name")
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Populate the book with its JSON data
	err := r.c.addBookWithBody(name, json)
	if err != nil {
		c.JSONP(interpretError(err), fmt.Sprintf("Error: %v", err))
		return
	}
	c.Status(http.StatusCreated)
}

// interpretError takes an error from a controller and returns a corresponding
// http status code.
func interpretError(e error) int {
	switch e {
	case ErrBookNotFound:
		return http.StatusNotFound
	case ErrDuplicateTitle:
		return http.StatusConflict
	case ErrEmptyName:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func main() {
	fmt.Println("Launching Bookworm CRUD Service...")

	// Instantiate a local MongoDB client
	// TODO: Put into database layer initialization, passing back the DB struct
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	// Error handling and deferred disconnection
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Get a collection from the local database
	bookColl := client.Database("local").Collection("active_books")

	// Some local test data to experiment with
	localBook := bookData{Title: "Luigi", Author: "Nintendo", Desc: "Ghosts in a mansion, get the vacuum cleaner!"}

	// Add the book to the collection
	result, err := bookColl.InsertOne(ctx, localBook)
	id := result.InsertedID
	fmt.Println("Inesrtion ID: ", id)

	// Initialize the custom router and controller
	// TODO: Initialize the controller with the DB struct
	control := controller{}
	route := router{c: control}

	// Create a gin router "r" with default middleware
	r := gin.Default()

	// Define a function to listen for a GET request on a book name
	r.GET("/book/:name", route.getBook)

	// Define a function to listen for a GET request for all books
	r.GET("/book", route.getAllBooks)

	// Define a function to listen for a POST request on a book name
	r.POST("/book/:name", route.addBook)

	// Define a function to listen for a POST request with a body on a book
	r.POST("/bookJSON/:name", route.addBookWithBody)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
