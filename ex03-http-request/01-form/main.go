package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type anything int

func (m anything) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("we are in serveHTTP")
	fmt.Println(req)
	tpl.ExecuteTemplate(w, "index.gohtml", req.PostForm)
	//tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tpl *template.Template

func init() {
	fmt.Println("now we are on init")
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d anything
	fmt.Println("now we are on main")
	http.ListenAndServe(":8080", d)
}
