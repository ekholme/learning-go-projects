package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

//defining data structures
type Page struct {
	Title string
	Body  []byte //note this is a byte rather than a string bc that's the type expected by the io libraries
}

//read this as 'this is a method named save that takes as its receiver p, a pointer to Page. It takes
//no parameters and returns a value of type error
//it returns an error bc it will return nil if everything goes well, and will only return anything if something goes wrong
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil { //this will return nil if we load a page and error otherwise
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil //this will return a pointer to a Page literal constructed with title and body values
}

//creating a template render function
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

//this lets us view a page
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	//if statement below will redirect a user to an edit page if the page doesn't exist
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

//editHandler will let us edit pages
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//saveHandler will handle the submission of forms located on edit pages
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//RESUME AT 'ERROR HANDLING' here: https://golang.org/doc/articles/wiki/

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
