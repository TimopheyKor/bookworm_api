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
