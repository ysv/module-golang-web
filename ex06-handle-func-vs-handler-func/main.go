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

	// We can pass this c as Handler interface because
	// HandlerFunc type implement ServeHttp as described in http package.
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}
