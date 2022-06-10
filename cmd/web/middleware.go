package main

import (
	"fmt"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request with method: %s hit path %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r) //call ServeHTTP to allow it go to next middle
	})
}
