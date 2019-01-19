package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, _ *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, _ *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	// Use DefaultServeMux if Handler is nil.
	http.ListenAndServe(":8080", nil)
}
