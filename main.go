package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
	//name := c.Param("name")
}

// interpretError takes an error from a controller and returns a corresponding
// http status code.
func interpretError(e error) int {
	switch e {
	case ErrBookNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func main() {
	fmt.Println("Launching Bookworm CRUD Service...")

	// Some local test data to experiment with
	bl := booksList{Books: make(map[string]bookData)}
	b0 := bookData{Title: "Luigi", Author: "Nintendo", Desc: "Ghosts in a mansion, get the vacuum cleaner!"}
	bl.Books["Luigi"] = b0

	// Initialize the custom router and controller
	control := controller{db: bl}
	route := router{c: control}

	// Create a gin router "r" with default middleware
	r := gin.Default()

	// Define a function to listen for a GET request on a book name
	r.GET("/book/:name", route.getBook)
	// Define a function to listen for a GET request for all books
	r.GET("/book", route.getAllBooks)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
