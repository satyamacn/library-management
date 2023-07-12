package main

import (
	"github.com/satyamacn/library-management/authors"
	"github.com/satyamacn/library-management/books"
	// "bufio"
	// "fmt"
	// "os"
	// "strings"
)


func main() {
	b := books.GetAllBooks()
	a := authors.GetAllAuthors()

	b.PrintAll()

	b.GetBookByID("11")
	a.GetAuthorByName("Jane Smith")
	// newBook := books.Book{
	// 	ID:        "11",
	// 	Title:     "ikigai",
	// 	Author:    authors.Author{Name: "Hector Garcia", Country: "USA", PenName: "H.G"},
	// 	Genre:     "Non-fiction",
	// 	Publisher: "XYZ",
	// 	Language:  "English",
	// }
	
	// books.AddBook(newBook)
	books.RemoveBookByName("ikigai")
	b.PrintAll();
    a.PrintAllAuthors();


	}

