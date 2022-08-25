package main

import "errors"

// Simple data structure for testing Gin & JSON marshalling.
// When using Mongo, this will instead use the Mongo driver queries and
// reference the database instead of using a local map.
type booksList struct {
	Books map[string]bookData `json:"books"`
}

// getBook takes a name and checks to see if the book exists inside the
// Books map, then returns either the bookData or an error.
func (bl *booksList) getBook(name string) (*bookData, error) {
	book, ok := bl.Books[name]
	if !ok {
		return nil, errors.New("nonexistent")
	}
	return &book, nil
}

// getAllBooks returns the entire Books map.
func (bl *booksList) getAllBooks() map[string]bookData {
	return bl.Books
}
