package handlers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
  "fmt"
  "log"
  "os"
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

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	_, err := mux.CurrentRoute(r).Subrouter().Get("view").URL()
  title := mux.Vars(r)["topic"]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, 302)
    return
	}

	renderTemplate(w, "view", p)
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
   if len(m) < 1 {
      return h
   }
   wrapped := h
   // loop in reverse to preserve middleware order
   for i := len(m) - 1; i >= 0; i-- {
      wrapped = m[i](wrapped)
   }
   return wrapped
}

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    log.SetOutput(os.Stdout) // logs go to Stderr by default
    log.Println(r.Method, r.URL)
    h.ServeHTTP(w, r) // call ServeHTTP on the original handler

  })
}

func HeartBeat(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "### alive ###")
}

