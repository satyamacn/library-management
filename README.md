# Library Management System

- Implemented a basic REST API using the Go `net/http` package. 
- The API has various endpoints that can be used to interact with the Library Management System data.

- Here are the endpoints that have been implemented:

1. **Get All Books** : `GET /api/books` - This endpoint retrieves all the books in the library.

2. **Add Book** : `POST /api/books` - This endpoint allows you to add a new book to the library. We need to send a JSON script containing the book details (ID, Title, Author, Genre, Publisher, Language) in the request body.

3. **Get Book by ID** : `GET /api/books/{id}` - This endpoint retrieves a specific book from the library based on its ID.

4. **Delete Book by Name** : `DELETE /api/books/{name}` - This endpoint deletes a book from the library based on its title (name).

5. **Get All Authors** : `GET /api/authors` - This endpoint retrieves all the authors from the library.


6. **Delete Last Book Record**: `DELETE /api/books/last` - This endpoint removes the last book record from the library.
