package main

import (
	"fmt"
	MI "github.com/fern4lvarez/go-metainspector/metainspector"
	"html/template"
	"net/http"
	"net/url"
	"os"
)

const lenMI = len("/mi/")

var templates = template.Must(template.ParseFiles("tmpl/index.html"))

type MetaInspector struct {
	Title string
}

func getPort() string {
	if p := os.Getenv("PORT"); p != "" {
		return p
	}
	return "8080"
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
	var ust string
	u, err := url.Parse(r.FormValue("url"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if u.Scheme != "" {
		u.Scheme = ""
		ust = u.String()[2:]
	} else {
		ust = u.String()
	}
	http.Redirect(w, r, "/mi/"+ust, http.StatusFound)
}

func miHandler(w http.ResponseWriter, r *http.Request) {
	uri := getUrl(w, r)
	fmt.Printf("Inspecting %s...", uri)
	mi, err := MI.New(uri)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(" Done: ", "http://"+r.Host+r.RequestURI)
	renderTemplate(w, "index", mi)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/inspect", inspectHandler)
	http.HandleFunc("/mi/", miHandler)

	port := getPort()
	fmt.Println("Listening to port", port)
	http.ListenAndServe(":"+port, nil)
}
