package main

import (
	"fmt"
	"net/http"

	"github.com/satyamacn/library-management/routers"
)

func main() {
	r := routers.SetupRouter()

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

