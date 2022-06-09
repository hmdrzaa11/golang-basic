package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

//map of some functions that we need to call inside of the template
var functions = template.FuncMap{}

// RenderTemplate uses the template cache to parse a template
func RenderTemplate(w http.ResponseWriter, templateName string) {
	tmpCache, err := createTemplateCache()
	if err != nil {
		//app should die
		log.Fatal(err)
	}
	if template, ok := tmpCache[templateName]; ok {
		template.Execute(w, nil)
	}
}

// createTemplateCache creates a map of template and their name as a cache
func createTemplateCache() (map[string]*template.Template, error) {
	//its going to build all template and combine them with the "base" layout at the startup of the App and keeps it in cache
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return nil, err
	}

	for _, pagePath := range pages {
		filename := filepath.Base(pagePath)
		ts, err := template.New(filename).Funcs(functions).ParseFiles(pagePath) //tell it to create a template based on "pagePath"
		//but without any layout yet
		if err != nil {
			return templateCache, err
		}
		//check to see if there is any layout file in templates dirs as well
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			//means we found at least 1 template
			//now we can parse that template with the layout
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return templateCache, err
			}
		}
		//now add the template to the cache
		templateCache[filename] = ts
	}
	return templateCache, nil
}
