package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

const (
	VIEWPATH  = "/view/"
	ADDPATH   = "/add/"
	INDEXPATH = "/"

	METAPATH = "meta.json"
	DATADIR  = "data"
)

var meta *MetaStore = LoadMetaStore(METAPATH)
var templates = template.Must(template.ParseFiles("add.html", "view.html", "index.html"))

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len(VIEWPATH):]
	data := LoadPageData(DATADIR, id)
	err := templates.ExecuteTemplate(w, "view.html", data)
	if err != nil {
		fmt.Fprintf(w, "Error loading view")
	}
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		err := templates.ExecuteTemplate(w, "add.html", nil)
		if err != nil {
			fmt.Fprintf(w, "Error loading add")
		}
	} else if r.Method == "POST" {
		inputText := r.FormValue("input")
		id := GetId(inputText)
		go RunSimulation(inputText, id)
		meta.AddAndWriteId(id)
		http.Redirect(w, r, "/view/"+id, http.StatusFound)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Fprintf(w, "Error loading index")
		return
	}
	t.Execute(w, meta)
}

func main() {
	// Make sure data directory exists
	_ = os.Mkdir(DATADIR, 0755)

	// Register handlers
	http.HandleFunc(VIEWPATH, viewHandler)
	http.HandleFunc(INDEXPATH, indexHandler)
	http.HandleFunc(ADDPATH, addHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve forever
	http.ListenAndServe(":8080", nil)
}
