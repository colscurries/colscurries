package main

import (
	"html/template"
	"net/http"
	"path"
)

func main() {
	fs := http.FileServer(http.Dir("assets/"))

	http.HandleFunc("/", index)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, fp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
