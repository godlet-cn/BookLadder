package entity

import (
	"errors"
	"strings"
	"time"
)

//BookStore stores all the books
type BookStore struct {
	Books []Book
}

//AddBook add a book to BookStore
func (store *BookStore) AddBook(book Book) (bool, error) {
	if book.ISBN == "" {
		return false, errors.New("a book's isbn is needed")
	}
	if store.IsExist(book.ISBN) {
		return false, errors.New("This book is existed in book store,with isbn: " + book.ISBN)
	}
	book.SaveTime = time.Now()
	store.Books = append(store.Books, book)
	return true, nil
}

//AddBooks add a list of books to BookStore
func (store *BookStore) AddBooks(books []Book) (bool, error) {
	for _, book := range books {
		if isSuccedd, err := store.AddBook(book); err != nil {
			return isSuccedd, err
		}
	}
	return true, nil
}

//IsExist check if BookStore has a book by give it's isbn
func (store BookStore) IsExist(isbn string) bool {
	for _, book := range store.Books {
		if book.ISBN == isbn {
			return true
		}
	}
	return false
}

//ListBooks add a list of books to BookStore
func (store *BookStore) ListBooks() []Book {
	return store.Books
}

//FindBookByIsbn find books by a given isbn
func (store *BookStore) FindBookByIsbn(isbn string) *Book {
	for index := range store.Books {
		if store.Books[index].ISBN == isbn {
			return &store.Books[index]
		}
	}
	return nil
}

//FindBooksByName find books by a given name
func (store *BookStore) FindBooksByName(name string) []Book {
	var result []Book
	for index := range store.Books {
		if strings.Contains(store.Books[index].Name, name) {
			result = append(result, store.Books[index])
		}
	}
	return result
}

//EditName modify a book's name
func (store *BookStore) EditName(isbn, name string) bool {
	for index := range store.Books {
		if store.Books[index].ISBN == isbn {
			store.Books[index].Name = name
			return true
		}
	}
	return false
}

//DeleteBook delete a book
func (store *BookStore) DeleteBook(isbn string) bool {
	for index, book := range store.Books {
		if book.ISBN == isbn {
			store.Books = append(store.Books[:index], store.Books[index+1:]...)
			return true
		}
	}
	return false
}
