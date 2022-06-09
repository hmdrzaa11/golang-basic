package main

import (
	"fmt"
	"net/http"

	"github.com/hmdrzaa11/hello-world/pkg/handlers"
)

const (
	portNumber = ":8080"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Listening on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
