package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
  "github.com/gorilla/mux"
)

type Page struct {
	Title string
	Body  []byte
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  t, err := template.ParseFiles("templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
  title := mux.Vars(r)["topic"]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func Edit(w http.ResponseWriter, r *http.Request) {
  title := mux.Vars(r)["topic"]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func View(w http.ResponseWriter, r *http.Request) {
  title := mux.Vars(r)["topic"]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, 302)
    return
	}

	renderTemplate(w, "view", p)
}
