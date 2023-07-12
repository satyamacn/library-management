package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/satyamacn/library-management/authors"
	"github.com/satyamacn/library-management/books"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Library Management System!")

	for {
		displayMenu()

		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			displayAllBooks()
		case "2":
			addBook(scanner)
		case "3":
			displayAllAuthors()
		case "4":
			removeBookByName(scanner)
		case "5":
			fmt.Println("Exiting the program...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func displayMenu() {
	fmt.Println("\nPlease select an option:")
	fmt.Println("1. Display all books")
	fmt.Println("2. Add a book")
	fmt.Println("3. Display all authors")
	fmt.Println("4. Remove book by name")
	fmt.Println("5. Exit")
	fmt.Print("Option: ")
}

func displayAllBooks() {
	fmt.Println("Displaying all books:")
	bookList := books.GetAllBooks()
	for _, book := range bookList.Books {
		fmt.Println("ID:", book.ID)
		fmt.Println("Title:", book.Title)
		fmt.Println("Author:", book.Author.Name)
		fmt.Println("Genre:", book.Genre)
		fmt.Println("Publisher:", book.Publisher)
		fmt.Println("Language:", book.Language)
		fmt.Println("----------------------")
	}
}

func addBook(scanner *bufio.Scanner) {
	fmt.Println("Adding a book:")

	fmt.Print("Enter the ID: ")
	scanner.Scan()
	id := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the title: ")
	scanner.Scan()
	title := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the author: ")
	scanner.Scan()
	author := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the genre: ")
	scanner.Scan()
	genre := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the publisher: ")
	scanner.Scan()
	publisher := strings.TrimSpace(scanner.Text())

	fmt.Print("Enter the language: ")
	scanner.Scan()
	language := strings.TrimSpace(scanner.Text())

	newBook := books.Book{
		ID:        id,
		Title:     title,
		Author:    authors.Author{Name: author},
		Genre:     genre,
		Publisher: publisher,
		Language:  language,
	}

	books.AddBook(newBook)
	fmt.Println("Book added successfully!")
}

func displayAllAuthors() {
	fmt.Println("Displaying all authors:")
	authorList := authors.GetAllAuthors()
	for _, author := range authorList.Authors {
		fmt.Println("ID:", author.ID)
		fmt.Println("Name:", author.Name)
		fmt.Println("Country:", author.Country)
		fmt.Println("Pen Name:", author.PenName)
		fmt.Println("----------------------")
	}
}

func removeBookByName(scanner *bufio.Scanner) {
	fmt.Println("Removing a book by name:")

	fmt.Print("Enter the book name: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	books.RemoveBookByName(name)
	fmt.Println("Book removed successfully!")
}