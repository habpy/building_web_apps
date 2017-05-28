package main

// Bill Gates: "If you want to be succesfull, get in front of what's coming and let it hit you"
// URI + URN = URL
import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type pageData struct {
	Title     string
	FirstName string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", idx)
	http.HandleFunc("/about", abot)
	http.HandleFunc("/contact", cntct)
	http.HandleFunc("/apply", aply)
	http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Index Page",
	}

	err := tpl.ExecuteTemplate(w, "index.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Kan index.html niet lezen", http.StatusInternalServerError)
		return
	}
}

func abot(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "About Page",
	}

	err := tpl.ExecuteTemplate(w, "about.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Kan about.html niet lezen", http.StatusInternalServerError)
		return
	}
}

func cntct(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Contact Page",
	}

	err := tpl.ExecuteTemplate(w, "contact.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Kan contact.html niet lezen", http.StatusInternalServerError)
		return
	}
}

func aply(w http.ResponseWriter, req *http.Request) {

	pd := pageData{
		Title: "Apply Page",
	}

	var first string
	if req.Method == http.MethodPost {
		first = req.FormValue("fname")
		pd.FirstName = first
	}

	err := tpl.ExecuteTemplate(w, "apply.html", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Kan apply.html niet lezen", http.StatusInternalServerError)
		return
	}
}
