package main

import (
	"errors"
	"net/http"
)

// Define global errors.
var (
	ErrBookNotFound   = errors.New("book not found in db")
	ErrDuplicateTitle = errors.New("book title already exists in db")
	ErrEmptyName      = errors.New("book title cannot be empty")
)

// interpretError takes an error from a controller and returns a corresponding
// http status code.
func interpretError(e error) int {
	switch e {
	case ErrBookNotFound:
		return http.StatusNotFound
	case ErrDuplicateTitle:
		return http.StatusConflict
	case ErrEmptyName:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

// panicIfError takes an error, and sends a panic on the error if it's not nil.
func panicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
