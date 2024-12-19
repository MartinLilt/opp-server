package main

import (
	"fmt"
	"go-backend-server/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	fmt.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}
