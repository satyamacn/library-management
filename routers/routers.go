package routers

import (
	"github.com/gorilla/mux"
	"github.com/satyamacn/library-management/handler"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
    r.HandleFunc("/api/books", handler.HandleBooks).Methods("GET", "POST")
	r.HandleFunc("/api/books/{id}", handler.HandleBookByID).Methods("GET")
	r.HandleFunc("/api/books/title/{title}", handler.HandleDeleteBookByTitle).Methods("DELETE")
	r.HandleFunc("/api/authors", handler.HandleAuthors).Methods("GET")
	r.HandleFunc("/api/authors/{name}", handler.HandleAuthorByName).Methods("GET")
	r.HandleFunc("/api/books/last", handler.HandleRemoveLastBook).Methods("DELETE")

	return r
}
