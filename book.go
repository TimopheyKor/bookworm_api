package main

// bookData is an object containing simple data representing a book.
type bookData struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}
