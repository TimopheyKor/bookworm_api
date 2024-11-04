package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// Router to access a controller and a logger if necessary. The router should
// handle all http & gin methods.
type router struct {
	c controller
}

// getBook takes the name of a book from the request context, calls to the
// controller to attempt to find a book with a title matching the name, and returns
// a JSON response or an error based on if the book exists in the dataset or not.
func (r *router) getBook(c *gin.Context) {
	name := c.Param("name")
	book, err := r.c.getBook(name, c)
	if err != nil {
		c.JSONP(interpretError(err), gin.H{"error": err.Error()})
		return
	}
	c.JSONP(http.StatusOK, book)
}

// getAllBooks returns a JSON response containing a map of all the books
// currently inside the database.
func (r *router) getAllBooks(c *gin.Context) {
	bookData, err := r.c.getAllBooks(c)
	if err != nil {
		c.JSONP(interpretError(err), gin.H{"error": err.Error()})
		return
	}

	c.JSONP(http.StatusOK, gin.H{"books": bookData})
}

// addBook takes a JSON-formatted body from a post request, attempts to bind the
// JSON data to a bookData object, then attempts to add it as a book document
// into the active_books collection of the database.
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

// TODO: Create a variant of deleteBook which returns the deleted book's info.
// deleteBook takes the name of a book from request context, then calls to the
// controller to attempt to find and delete a book with that name from the
// database.
func (r *router) deleteBook(c *gin.Context) {
	name := c.Param("name")
	err := r.c.deleteBook(name, c)
	if err != nil {
		c.JSONP(interpretError(err), gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

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
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := database.disconnect(clientContext); err != nil {
			// slog.Error("disconnect db", slog.Any("error", err))
			fmt.Fprintf(os.Stderr, "disconnect db: %s", err)
			os.Exit(1)
		}
	}()

	// Initialize the custom router and controller
	control := controller{db: database}
	route := router{c: control}

	// Create a gin router "r" with default middleware
	r := gin.Default()

	// TODO: Make all the listener functions interact with the database.
	// Listen for a GET request on a book name, which returns a book of that name.
	r.GET("/books/:name", route.getBook)

	// Listen for a GET request, which returns all the books in the database.
	r.GET("/books", route.getAllBooks)

	// Listen for a POST request, which adds a new book to the database.
	r.POST("/books", route.addBook)

	// Listen for a DELETE request on a book name, which attempts to delete a
	// book of that name.
	r.DELETE("/books/:name", route.deleteBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
