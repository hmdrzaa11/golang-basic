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
	var app config.AppConfig
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}
	app.TemplateCache = templateCache
	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Listening on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
