package main

import (
	"context"
)

// controller holds a reference to the database and has methods which call
// database queries, then format the responses into an appropriate state
// for the router to handle.
type controller struct {
	db *mongoDB
}

// addBook gets a bookData object and a request context from the router, then
// calls to the databse layer to add that book to the database.
func (c *controller) addBook(data bookData, ctx context.Context) error {
	err := c.db.addBook(data, ctx)
	return err
}

// getAllBooks gets a request context from the router and calls to the database
// layer to retrieve all of the books data, which it then returns to the router.
func (c *controller) getAllBooks(ctx context.Context) ([]bookData, error) {
	allBooks, err := c.db.getAllBooks(ctx)
	if err != nil {
		return nil, err
	}
	return allBooks, nil
}

// TODO: Figure out what functionality can be added to this controller function.
// getBook takes a name string and request context, calls to the database layer
// to find a book with a matching title, then returns the book's data to the
// router if it's successful.
func (c *controller) getBook(name string, ctx context.Context) (*bookData, error) {
	book, err := c.db.getBook(name, ctx)
	return book, err
}

// TODO: Figure out what functionality can be added to this controller function.
// deleteBook takes a name string and a request context, calls to the database
// layer to find and delete the entry of the first book who's title matches, and
// returns an error if unsuccessful.
func (c *controller) deleteBook(name string, ctx context.Context) error {
	err := c.db.deleteBook(name, ctx)
	return err
}
