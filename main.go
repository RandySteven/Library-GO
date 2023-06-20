package main

import (
	"log"
	"math/rand"
)

type Book struct {
	ID         int64
	Title      string
	Contents   map[int]string
	LatestPage int
}

type Books []Book

type Library struct {
	Books      Books
	ActiveBook int64
}

type IBookAction interface {
	AddBook(title string, contents map[int]string) *Book
	GetAllBooks(books Books)
	GetLatestPage(book Book) int
	StoreBooksToLibrary(books Books) *Library
	ReadContent(readBook *Book, readContents int) *Book
}

func (book *Book) AddBook(title string, contents map[int]string) *Book {

	book = &Book{
		ID:         int64(rand.Uint64()),
		Title:      title,
		Contents:   contents,
		LatestPage: 0,
	}

	return book
}

func (book *Book) ReadContent(readBook *Book, readContents int) *Book {
	page := -1
	if readContents > len(readBook.Contents) {
		log.Printf("There is no page %d for %s", readContents, readBook.Title)
	} else {
		for k := range readBook.Contents {
			if readContents == k {
				page = k
			}
		}
		readBook.LatestPage = page
	}

	return readBook
}

func (book *Book) GetAllBooks(books Books) {
	for i := 0; i < len(books); i++ {
		log.Println("====================================================")
		log.Println("ID : ", books[i].ID)
		log.Printf("Title of books = %s", books[i].Title)
		for page, content := range books[i].Contents {
			log.Printf("Page = %d", page)
			log.Printf("Content = %s", content)
		}
		log.Println("====================================================")
	}
}

func (book *Book) StoreBooksToLibrary(books Books) *Library {
	library := &Library{
		Books: books,
	}
	return library
}

type ILibraryAction interface {
	ReadBook(bookID int64) int64
	ReadBookByTitle(bookTitle string) int64
	GetAllBooks()
	GetBook(bookId int64) Book
}

func (library *Library) GetBook(bookID int64) int64 {
	library.ActiveBook = bookID
	return library.ActiveBook
}

func (library *Library) GetAllBooks() {
	for i := 0; i < len(library.Books); i++ {
		log.Println("====================================================")
		log.Println("ID : ", library.Books[i].ID)
		log.Printf("Title of books = %s", library.Books[i].Title)
		for page, content := range library.Books[i].Contents {
			log.Printf("Page = %d", page)
			log.Printf("Content = %s", content)
		}
		log.Println("====================================================")
	}
}

func (library *Library) GetBookByTitle(title string) int64 {
	index := -1
	search := false
	for i := 0; i < len(library.Books); i++ {
		if library.Books[i].Title == title {
			index = i
			search = true
		}
	}

	if search == true {
		library.ActiveBook = library.Books[index].ID
	}

	return library.ActiveBook
}

func (library *Library) ReadBook(bookID int64) Book {
	search := false
	index := -1
	for i := 0; i < len(library.Books); i++ {
		if library.Books[i].ID == bookID {
			index = i
			search = true
		}
	}

	if search == true {
		return library.Books[index]
	}

	return Book{}
}

func main() {
	var books Books

	title1 := "Title 1"
	contents1 := make(map[int]string)

	contents1[1] = "Test"
	contents1[2] = "Make"

	book := &Book{}
	book1 := book.AddBook(title1, contents1)
	log.Println(book1)

	title2 := "Title 2"
	contents2 := make(map[int]string)

	contents2[1] = "Alo alo"
	contents2[2] = "Babi"

	book2 := book.AddBook(title2, contents2)
	log.Println(book2)

	books = append(books, *book1)
	books = append(books, *book2)

	library := book.StoreBooksToLibrary(books)
	library.GetAllBooks()

	readBook := library.GetBookByTitle("Title 2")
	log.Printf("Now you read book : %s", library.ReadBook(readBook).Title)
}
