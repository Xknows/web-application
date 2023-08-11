package main

import (
	"fmt"
	"html/template"
	"net/http"
)
var port = ":8889" 

func Home(rw http.ResponseWriter, r *http.Request){
	renderTemplate(rw, "home.page.tmpl")
}

func about(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintf(rw,"yep!")
}


func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("home.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}


func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", about)
	
	fmt.Println("running on port:",port)
	_ = http.ListenAndServe(port, nil)
	
}



