package main

import (
	"fmt"
	"net/http"
	"web/packeges/handlers"
)

const port = ":8889"

// main function
func main() {
	http.HandleFunc("/home", handlers.Home)

	fmt.Println("running on port:", port)
	_ = http.ListenAndServe(port, nil)

}
