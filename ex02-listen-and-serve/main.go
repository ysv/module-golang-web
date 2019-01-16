package main

import (
	"fmt"
	"net/http"
)

type anything int

func (a anything) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Anything type rulez\n")
}

func main() {
	var a anything
	http.ListenAndServe(":8080", a)
}
