package main

import (
	MI "github.com/fern4lvarez/go-metainspector/metainspector"
	"html/template"
	"net/http"
	"os"
)

const lenMI = len("/mi/")

var templates = template.Must(template.ParseFiles("tmpl/index.html"))

type MetaInspector struct {
	Title string
}

func getUrl(w http.ResponseWriter, r *http.Request) string {
	return r.URL.Path[lenMI:]
}

func renderEmptyTemplate(w http.ResponseWriter, tmpl string) {
	err := templates.ExecuteTemplate(w, tmpl+".html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, mi *MI.MetaInspector) {
	err := templates.ExecuteTemplate(w, tmpl+".html", mi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderEmptyTemplate(w, "index")
}

func inspectHandler(w http.ResponseWriter, r *http.Request) {
  url := r.FormValue("url")
  http.Redirect(w, r, "/mi/"+url, http.StatusFound)
}


func miHandler(w http.ResponseWriter, r *http.Request) {
	uri := getUrl(w, r)
	mi, err := MI.New(uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "index", mi)
	//http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/inspect", inspectHandler)
	http.HandleFunc("/mi/", miHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
