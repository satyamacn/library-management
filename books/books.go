package books

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/satyamacn/library-management/authors"
)

const fileName = "books.json"

type Book struct {
	ID        string
	Title     string
	Author    authors.Author
	Genre     string
	Publisher string
	Language  string
}

type Books struct {
	Books []Book
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

func (b Books) PrintAll() {
	for i := 0; i < len(b.Books); i++ {
		fmt.Println(b.Books[i].ToString())
	}
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

func RemoveBookByName(title string) {
	books := GetAllBooks()

	index := -1
	for i, book := range books.Books {
		if book.Title == title {
			index = i
			break
		}
	}

	if index != -1 {
		books.Books = append(books.Books[:index], books.Books[index+1:]...)
	} else {
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

	log.Printf("Book '%s' removed successfully.", title)
}

func RemoveLastBook() {
	books := GetAllBooks()

	if len(books.Books) == 0 {
		log.Println("No books in the library.")
		return
	}

	books.Books = books.Books[:len(books.Books)-1]

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

	log.Println("Last book record removed successfully.")
}

