package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()

	// This will handle dog/something/or/someone
	mux.Handle("/dog/", d)
	// This will handle only exact match.
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
