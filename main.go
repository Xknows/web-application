package main

import (
	"fmt"
	"net/http"
)
var port = ":8889" 

func Home(wr http.ResponseWriter, r *http.Request){
	fmt.Fprintf(wr,"Home page")
}

func about(wr http.ResponseWriter, r *http.Request){
	fmt.Fprintf(wr,"yep!")
}


func main() {
	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", about)
	
	fmt.Println("running on port:",port)
	_ = http.ListenAndServe(port, nil)
	
}




