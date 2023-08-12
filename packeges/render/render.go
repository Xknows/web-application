package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// render a specific template

func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	// to reduce reading from disk, cache is used
	// create
	templateCache, error := createTemplateCache()
	if error != nil {
		log.Fatal(error)
	}
	// use and check for true
	template, check := templateCache[tmpl]
	if !check {
		log.Fatal(error)
	}
	// execute elemnt of map
	buffer := new(bytes.Buffer)
	error = template.Execute(buffer, nil)
	if error != nil {
		log.Fatal(error)
	}

	// render

	_, error = buffer.WriteTo(rw)
	if error != nil {
		log.Fatal(error)
	}
}

// a map of string pointing to "*template.Template" for storing templates
func createTemplateCache() (map[string]*template.Template, error) {
	myChace := make(map[string]*template.Template)
	// get all *.page.tmpl files
	pages, error := filepath.Glob("./templates/*.page.tmpl")
	if error != nil {
		return myChace, error
	}
	// stor all to "name" of files by real name not path name
	for _, page := range pages {
		name := filepath.Base(page)
		// parse
		templateSet, error := template.New(name).ParseFiles(page)
		if error != nil {
			return myChace, error
		}
		// for layout
		layouts, error := filepath.Glob("./templates/*.layout.tmpl")
		if error != nil {
			return myChace, error
		}
		//parse each layout may template need that
		if len(layouts) > 0 {
			templateSet, error = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if error != nil {
				return myChace, error
			}
		}
		// key of map equel to templateset
		myChace[name] = templateSet
	}
	return myChace, nil
}
