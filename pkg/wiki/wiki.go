package wiki

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

// Page for web page
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	_, currentFilePath, _, _ := runtime.Caller(0)
	filename := p.Title + ".txt"
	targetFilePath := filepath.Join(filepath.Dir(currentFilePath), filename)
	return os.WriteFile(targetFilePath, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {

	_, currentFilePath, _, _ := runtime.Caller(0)
	filename := title + ".txt"
	targetFilepath := filepath.Join(filepath.Dir(currentFilePath), filename)
	body, err := os.ReadFile(targetFilepath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
*/

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler starting")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	_, currentFilePath, _, _ := runtime.Caller(0)
	editHtmlpath := filepath.Join(filepath.Dir(currentFilePath), "edit.html")
	viewHtmlpath := filepath.Join(filepath.Dir(currentFilePath), "view.html")

	var templates = template.Must(template.ParseFiles(editHtmlpath, viewHtmlpath))
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ViewHandler starting")
	title := r.URL.Path[len("/view/"):]
	p, err := LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("EditHandler starting")
	title := r.URL.Path[len("/edit/"):]
	p, err := LoadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("SaveHandler starting")
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
