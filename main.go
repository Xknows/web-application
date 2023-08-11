package main

import (
	"fmt"
	"net/http"
)


var port = ":8888" 

func main() {

	http.HandleFunc("/", func(wr http.ResponseWriter, r *http.Request ){
		_,_ = fmt.Fprintf(wr, "hello world")
	

	})
 _ = http.ListenAndServe(port, nil)
}




