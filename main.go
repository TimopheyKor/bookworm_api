package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Simple structures for testing Gin & JSON marshalling
type bookData struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}
type booksList struct {
	Books map[string]bookData `json:"books"`
}

// getBook takes the name of a book from the request context, and returns
// a JSON response or an error based on if the book exists in the dataset or not.
func (bl *booksList) getBook(c *gin.Context) {
	name := c.Param("name")
	book, ok := bl.Books[name]
	if ok {
		c.JSONP(http.StatusOK, book)
		return
	}
	c.Status(http.StatusNotFound)
}

func main() {
	fmt.Println("Launching Bookworm CRUD Service...")

	// Some local test data to experiment with
	bl := booksList{Books: make(map[string]bookData)}
	b0 := bookData{Title: "Luigi", Author: "Nintendo", Desc: "Ghosts in a mansion, get the vacuum cleaner!"}
	bl.Books["Luigi"] = b0

	// Create a gin router "r" with default middleware
	r := gin.Default()

	// Define a function to listen for a GET request on a book name
	r.GET("/book/:name", bl.getBook)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
