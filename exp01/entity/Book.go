package entity

import (
	"fmt"
	"time"
)

//BookType which type of book
type BookType int

const (
	Science BookType = iota
	Art
)

//Book represent a book information
type Book struct {
	ISBN      string
	Name      string
	Author    string
	Publisher string
	Type      BookType
	SaveTime  time.Time
}

//String print book information in a formal way
func (book Book) String() string {
	return fmt.Sprintf("<<%T Isbn:%s Name:%s Author:%s Publisher:%s >>", book, book.ISBN, book.Name, book.Author, book.Publisher)
}
