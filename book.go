package main

// bookData is an object containing simple data representing a book.
type bookData struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"description"`
}

// NewBookData takes a name string and returns a book with the title field
// filled with the name.
func NewBookData(name string) bookData {
	return bookData{Title: name}
}
