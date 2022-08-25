package main

import "errors"

// controller holds a reference to the database and has methods which call
// database queries, then format the responses into an appropriate state
// for the router to handle.
type controller struct {
	db booksList
}

// Define global errors.
var (
	ErrBookNotFound = errors.New("book not found in db")
)

// getBook calls the db's function to get a book, then returns either the
// bookData or an error based on if the book was found or not.
func (c *controller) getBook(name string) (*bookData, error) {
	book, err := c.db.getBook(name)
	if err != nil {
		return nil, ErrBookNotFound
	}
	return book, nil
}

// getAllBooks calls the db's function to get the entire map of books,
// and returns the map.
func (c *controller) getAllBooks() map[string]bookData {
	return c.db.getAllBooks()
}
