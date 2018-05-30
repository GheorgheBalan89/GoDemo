package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Hello you've requested: %s \n", r.URL.Path)
		indexData := renderIndex()
		templ := template.Must(template.ParseFiles("views/layout.html"))

		templ.Execute(w, indexData)

	})

	r.HandleFunc("/forms", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("views/forms.html"))
		if r.Method != http.MethodPost {
			templ.Execute(w, nil)
		} else {
			details := ContactDetails{
				Email:   r.FormValue("email"),
				Subject: r.FormValue("subject"),
				Message: r.FormValue("message"),
			}

			// do something with details
			fmt.Fprintf(w, "Email %s, subject %s, message %s", details.Email, details.Subject, details.Message)
		}
		// tmpl.Execute(w, struct{ Success bool }{true})

	})

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s \n", title, page)
	})

	//***serve static assets
	// fs := http.FileServer(http.Dir("assets/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))
	// http.ListenAndServe(":8081", nil) ***

	http.ListenAndServe(":8081", r)

}

func renderIndex() TodoPageData {
	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Make a template", Done: false},
			{Title: "Add sass", Done: false},
			{Title: "Add authentication", Done: false},
		},
	}

	return data
}
