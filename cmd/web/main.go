package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/hmdrzaa11/hello-world/pkg/config"
	"github.com/hmdrzaa11/hello-world/pkg/handlers"
	"github.com/hmdrzaa11/hello-world/pkg/render"
)

const (
	portNumber = ":8080"
)

var (
	app     config.AppConfig
	session *scs.SessionManager
)

func main() {
	app.InProduction = false //set it to true when in production
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	//we are going to use the default "cookie" as store right now
	session.Cookie.Persist = true                  //after user closes the webpage the cookie persist
	session.Cookie.SameSite = http.SameSiteLaxMode //the restriction around what site this cookie need to apply to
	session.Cookie.Secure = app.InProduction       //needs a https connection

	//we assign the session to the app config
	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can not create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false           //do not use cache in development and build templates each time
	render.NewTemplate(&app)       //share appConfig to the render template
	repo := handlers.NewRepo(&app) //share appConfig to create a repo inside handlers
	handlers.NewHandlers(repo)

	fmt.Printf("Listening on port %s \n", portNumber)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Fatal(srv.ListenAndServe())
}
