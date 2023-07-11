package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the page handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the home page")
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

// About is the about page handler
func About(writer http.ResponseWriter, request *http.Request) {
	sum := addValues(2, 2)

	_, _ = fmt.Fprintf(writer, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}
