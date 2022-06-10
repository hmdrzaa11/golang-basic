package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("request with method: %s hit path %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r) //call ServeHTTP to allow it go to next middle
	})
}

// CsrfMiddleware adds CSRF protection to all the post requests
func CsrfMiddleware(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{ //use cookie to manage the csrf
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad this fn is going to load session on each request
func SessionLoad(next http.Handler) http.Handler {
	//we have a session already configured but now we need to use it
	return session.LoadAndSave(next) //because  "session" defined at the package level we can access it in here
}
