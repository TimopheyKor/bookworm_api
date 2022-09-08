package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Router to access a controller and a logger if necessary. The router should
// handle all http & gin methods.
type router struct {
	c controller
}

// getBook takes the name of a book from the request context, and returns
// a JSON response or an error based on if the book exists in the dataset or not.
/*
func (r *router) getBook(c *gin.Context) {
	name := c.Param("name")
	book, err := r.c.getBook(name)
	if err != nil {
		c.Status(interpretError(err))
		return
	}
	c.JSONP(http.StatusOK, book)
}
*/

// TODO: Modify getAllBooks for mongoDB controller
// getAllBooks returns a JSON response containing a map of all the books
// currently inside the database.
/*
func (r *router) getAllBooks(c *gin.Context) {
	c.JSONP(http.StatusOK, r.c.getAllBooks())
}
*/

// addBook takes a JSON-formatted body from a post request, and attempts to
// add it as a book document into the active_books database.
func (r *router) addBook(c *gin.Context) {
	var data bookData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := r.c.addBook(data, c)
	if err != nil {
		c.JSONP(interpretError(err), fmt.Sprintf("Error: %v", err))
		return
	}
	c.Status(http.StatusCreated)
}

// TODO: Modify addBookWithBody for mongoDB controller
// addBookWithBody checks to see if a JSON body has been given to the request,
// then checks if the body is valid. If so, it calls addBook to make a new
// entry into the database, then populates the book with the JSON body data.
/*
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
*/

func main() {
	fmt.Println("Launching Bookworm CRUD Service...")

	// TODO: Make the URI use environment variables.
	// Constants
	const (
		localURI = "mongodb://localhost:27017"
	)

	// Instantiate a local MongoDB client with a fresh context. Time out the
	// request if it's not completed in 10 seconds.
	clientContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	database, err := newMongoDB(clientContext, localURI)

	// Error handling and deferred disconnection
	panicIfError(err)
	defer func() {
		err = database.disconnect(clientContext)
		panicIfError(err)
	}()

	// Some local test data to experiment with
	localBook := bookData{Title: "Luigi", Author: "Nintendo", Desc: "Ghosts in a mansion, get the vacuum cleaner!"}

	// Initialize the custom router and controller
	control := controller{db: database}
	route := router{c: control}

	// TESTING: Add the localBook to the database through the controller.
	control.addBook(localBook, clientContext)

	// Create a gin router "r" with default middleware
	r := gin.Default()

	// TODO: Make  all the listener functions interact with the database.
	// Define a function to listen for a GET request on a book name
	//r.GET("/book/:name", route.getBook)

	// Define a function to listen for a GET request for all books
	//r.GET("/books", route.getAllBooks)

	// Define a function to listen for a POST request on a book name
	r.POST("/books/addBookData", route.addBook)

	// Define a function to listen for a POST request with a body on a book
	//r.POST("/bookJSON/:name", route.addBookWithBody)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
