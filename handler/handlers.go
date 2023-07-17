package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satyamacn/library-management/authors"
	"github.com/satyamacn/library-management/books"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		bookList := books.GetAllBooks()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(bookList)
	} else if r.Method == "POST" {
		var book books.Book
		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		books.AddBook(book)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Book added successfully!")
	}
}

func HandleBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	book := books.GetAllBooks().GetBookByID(id)
	if book.ID == "" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func HandleAuthors(w http.ResponseWriter, r *http.Request) {
	authorList := authors.GetAllAuthors()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authorList)
}

func HandleAuthorByName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	author := authors.GetAllAuthors().GetAuthorByName(name)
	if author.Name == "" {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func HandleDeleteBookByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	if r.Method == "DELETE" {
		books.RemoveBookByName(title)
		fmt.Fprintf(w, "Book with title '%s' removed successfully!", title)
	}
}

func HandleRemoveLastBook(w http.ResponseWriter, r *http.Request) {
	books.RemoveLastBook()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Last book record removed."))
}
