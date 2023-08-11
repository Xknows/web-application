package main

import (
	"fmt"
	"net/http"
	"web/packeges/handlers"
)

var port = ":8889"

func main() {
	http.HandleFunc("/home", handlers.Home)

	fmt.Println("running on port:", port)
	_ = http.ListenAndServe(port, nil)

}
