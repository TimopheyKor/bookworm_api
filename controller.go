package main

import "context"

// controller holds a reference to the database and has methods which call
// database queries, then format the responses into an appropriate state
// for the router to handle.
type controller struct {
	db *mongoDB
}

func (c *controller) addBook(data bookData, ctx context.Context) error {
	err := c.db.addBook(data, ctx)
	return err
}

// TODO: Write getAllBooks function
func (c *controller) getAllBooks() bookData {
	return bookData{}
}
