package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	fileName    = "books.json"
	authorsFile = "authors.json"
)

type Book struct {
	ID        string
	Title     string
	Author    Author
	Genre     string
	Publisher string
	Language  string
}

type Books struct {
	Books []Book
}

type Author struct {
	ID       string
	Name     string
	Country  string
	PenName  string
}

type Authors struct {
	Authors []Author
}

func CheckFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetAllBooks() Books {
	err := CheckFile(fileName)
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", fileName)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var books Books

	json.Unmarshal(byteValue, &books)

	return books
}

func GetAllAuthors() Authors {
	err := CheckFile(authorsFile)
	if err != nil {
		log.Println(err)
	}

	jsonFile, err := os.Open(authorsFile)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", authorsFile)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var authors Authors

	json.Unmarshal(byteValue, &authors)

	return authors
}

func (b Books) PrintAll() {
	for i := 0; i < len(b.Books); i++ {
		fmt.Println(b.Books[i].ToString())
	}
}

func (b Authors) PrintAllAuthors() {
	for i := 0; i < len(b.Authors); i++ {
		fmt.Println(b.Authors[i].ToString())
	}
}

func (b Author) ToString() string {
	return "Name: " + b.Name + ", " + "Pen Name: " + b.PenName + ", " + "Country: " + b.Country
}

func (b Book) ToString() string {
	return "ID: " + b.ID + ", " + "Title: " + b.Title + ", " + "Author: " + b.Author.Name + ", " + "Genre: " + b.Genre + ", " + "Publisher: " + b.Publisher + ", " + "Language: " + b.Language
}

func (b Books) GetBookByID(id string) Book {
	for i := 0; i < len(b.Books); i++ {
		if b.Books[i].ID == id {
			fmt.Printf("Book found: ID: %s\n Title: %s\n Author: %s\n Genre: %s\n Publisher: %s\n", b.Books[i].ID, b.Books[i].Title, b.Books[i].Author.Name, b.Books[i].Genre, b.Books[i].Publisher)
			return b.Books[i]
		}
	}

	return Book{}
}

func (a Authors) GetAuthorByName(name string) Author {
	for i := 0; i < len(a.Authors); i++ {
		if a.Authors[i].Name == name {
			fmt.Printf("Author found: Name: %s\n Country: %s\n Pen Name: %s\n", a.Authors[i].Name, a.Authors[i].Country, a.Authors[i].PenName)
			return a.Authors[i]
		}
	}

	return Author{}
}

func AddBook(book Book) {
	books := GetAllBooks()
	books.Books = append(books.Books, book)

	updatedBooksJSON, err := json.Marshal(books)
	if err != nil {
		log.Println("Error marshaling books:", err)
		return
	}

	err = ioutil.WriteFile(fileName, updatedBooksJSON, 0644)
	if err != nil {
		log.Println("Error writing books file:", err)
		return
	}

	fmt.Println("Book added successfully.")
}

func RemoveBookByName(name string) {
	books := GetAllBooks()

	// Find the index of the book with the given name
	index := -1
	for i, book := range books.Books {
		if book.Title == name {
			index = i
			break
		}
	}

	// If the book was found, remove it from the list
	if index != -1 {
		books.Books = append(books.Books[:index], books.Books[index+1:]...)
	} else {
		fmt.Println("Book not found.")
		return
	}

	updatedBooksJSON, err := json.Marshal(books)
	if err != nil {
		log.Println("Error marshaling books:", err)
		return
	}

	err = ioutil.WriteFile(fileName, updatedBooksJSON, 0644)
	if err != nil {
		log.Println("Error writing books file:", err)
		return
	}

	fmt.Println("Book removed successfully.")
}

func main() {

	books := GetAllBooks()
	authors := GetAllAuthors()

	books.PrintAll()

	books.GetBookByID("03")
	authors.GetAuthorByName("Jane Smith")
	newBook := Book{
		ID:        "04",
		Title:     "ikigai",
		Author:    Author{Name: "Hector Garcia", Country: "USA", PenName: "H.G"},
		Genre:     "Non-fiction",
		Publisher: "XYZ",
		Language:  "English",
	}

	AddBook(newBook)
	//RemoveBookByName("New Book")
	books.PrintAll();
	authors.PrintAllAuthors();
}