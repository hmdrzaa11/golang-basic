package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hmdrzaa11/hello-world/pkg/config"
	"github.com/hmdrzaa11/hello-world/pkg/handlers"
	"github.com/hmdrzaa11/hello-world/pkg/render"
)

const (
	portNumber = ":8080"
)

func main() {
	var app config.AppConfig //we are going to share this through entire app
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false           //do not use cache in development and build templates each time
	render.NewTemplate(&app)       //share appConfig to the render template
	repo := handlers.NewRepo(&app) //share appConfig to create a repo inside handlers
	handlers.NewHandlers(repo)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Listening on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
