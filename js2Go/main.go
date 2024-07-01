package main

import (
	"log"
	"net/http"
	"text/template"
)

type Todo struct {
	Id      int
	Message string
}

var data = map[string][]Todo{
	"Todos": {
		{Id: 1, Message: "Buy Milk"},
	},
}

func main() {

	todosHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("index.html"))
		if err := templ.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	addTodoHandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		if message != "" {
			todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
			data["Todos"] = append(data["Todos"], todo)

			// Return only the new item as an HTML fragment
			w.Header().Set("Content-Type", "text/html")
			templ := template.Must(template.New("todo").Parse(`<li>{{.Message}}</li>`))
			if err := templ.Execute(w, todo); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	http.HandleFunc("/", todosHandler)
	http.HandleFunc("/add-todo", addTodoHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
