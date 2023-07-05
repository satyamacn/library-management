package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const fileName string = "books.json"

type Books struct {
	Books []Book
}

type Book struct {
	ID        string
	Title     string
	Author    string
	Genre     string
	Publisher string
	Language  string
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
	return "ID: " + b.ID + ", " + "Title: " + b.Title + ", " + "Author: " + b.Author + ", " + "Genre: " + b.Genre + ", " + "Publisher: " + b.Publisher + ", " + "Language: " + b.Language
}

func (b Books) GetBookByID(id string) Book {
	for i := 0; i < len(b.Books); i++ {
		if b.Books[i].ID == id {
			fmt.Printf("Book found: ID: %s\n Title: %s\n Author: %s\n Genre: %s\n Publisher: %s\n", b.Books[i].ID, b.Books[i].Title, b.Books[i].Author,b.Books[i].Genre, b.Books[i].Publisher)
			return b.Books[i]
		}
	}
	
	return Book{}
}

func main() {
	books := GetAllBooks()

	books.PrintAll()

    books.GetBookByID("01")
}
