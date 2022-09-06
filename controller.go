package main

import (
	"errors"
)

// controller holds a reference to the database and has methods which call
// database queries, then format the responses into an appropriate state
// for the router to handle.
type controller struct {
	db booksList
}

// Define global messages.
var (
	AddNewSuccess = "sucessfully added new book to db"
)

// Define global errors.
var (
	ErrBookNotFound   = errors.New("book not found in db")
	ErrDuplicateTitle = errors.New("book title already exists in db")
	ErrEmptyName      = errors.New("book title cannot be empty")
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

// addNewBook takes a name string and calls the db's addNewBook function
// to add a book to the database. If sucessful, it returns a success message,
// otherwise it returns an error message and error code.
func (c *controller) addNewBook(name string) error {
	if name == "" {
		return ErrEmptyName
	}
	return c.db.addNewBook(name)
}

// addBookWithBody takes a book name and a json bookData object, then attempts
// to create a new entry in the database by calling addNewBook, then populating
// the value of that book name entry. It returns an error if addNewBook fails.
func (c *controller) addBookWithBody(name string, json bookData) error {
	err := c.addNewBook(name)
	c.db.populateBookData(name, json)
	return err
}
