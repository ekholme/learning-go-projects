package main

import (
	"fmt"
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

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}

//resume at 'Introducing the net/http package' in this doc: https://golang.org/doc/articles/wiki/
