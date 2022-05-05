package main

import (
	"net/http"
	"html/template"
	// "io"
	// "log"
)

type Blog struct {
	Title string
	Date int
	Body string

}

type Form struct {
	firstname string
	lastname string
	email string
	number int
}

func main() {
	fileserver := http.FileServer(http.Dir("../template/*.html"))
	http.Handle("/", fileserver)
	http.HandleFunc("/index.html", index)
	http.ListenAndServe(":2020", nil)
}

var tpl *template.Template

// func init () {
// 	tpl = template.Must(template.ParseGlob("../template/*.html"))
// }

func index(w http.ResponseWriter, r *http.Request) {
	// tpl.ExecuteTemplate(w, "index.html", nil)
	if r.url.Path != "/index.html" {
		http.Error(w, "404 not found".http.StatusNotFound).Method("POST")
		return
	}

	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}