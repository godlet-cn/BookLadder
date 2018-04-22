package main

import (
	"fmt"

	"github.com/godlet-cn/BookLadder/exp01/entity"
)

func main() {
	var bookStore entity.BookStore

	bookStore.AddBooks([]entity.Book{
		entity.Book{ISBN: "b3aer4ha", Name: "Gene Tecknology", Author: "Jack", Type: entity.Science, Publisher: "AF"},
		entity.Book{ISBN: "je3f3fga", Name: "Deep into Painting", Author: "Alice", Type: entity.Art, Publisher: "AF"},
		entity.Book{ISBN: "dei3w3f2", Name: "A journey to China", Author: "Wade", Type: entity.Science, Publisher: "BF"},
		entity.Book{ISBN: "l38de398", Name: "A new research or Black hole ", Author: "Jack", Type: entity.Science, Publisher: "CF"},
		entity.Book{ISBN: "ade3jk3d", Name: "Deep into photoshop", Author: "Lee", Type: entity.Art, Publisher: "CF"},
	})

	fmt.Println("List all the books")
	books := bookStore.ListBooks()
	for _, book := range books {
		fmt.Println(book)
	}
	fmt.Println()

	var isbn = "dei3w3f2"
	fmt.Println("Find book.isbn:", isbn)
	book := bookStore.FindBookByIsbn(isbn)
	fmt.Println(book)
	fmt.Println()

	fmt.Println("Change the book's author")
	book.Author = "Liu"
	book = bookStore.FindBookByIsbn(isbn)
	fmt.Println("The author's name is", book.Author)

	fmt.Println("Change the book's name")
	bookStore.EditName(isbn, "I Love China")
	book = bookStore.FindBookByIsbn(isbn)
	fmt.Println("Now the book's name is", book.Name)
	fmt.Println()

	fmt.Println("Find books by name.keyword:", "Deep into")
	books = bookStore.FindBooksByName("Deep into")
	for _, book := range books {
		fmt.Println(book)
	}
	fmt.Println()

	fmt.Println("Delete book. isbn:", isbn)
	bookStore.DeleteBook(isbn)
	book = bookStore.FindBookByIsbn(isbn)
	if book == nil {
		fmt.Println("Faild to find a book which's isbn is", isbn)
	} else {
		fmt.Println("Find the book:", book)
	}
}
