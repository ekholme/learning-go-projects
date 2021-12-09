package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

//will ensure that our template can compile when the program is initiated
//note that init() gets run when the program starts
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl)) //template.Must requires the template to compile, otherwise the program will fail
}

//setting the html template as a variable
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
</head>

<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
    <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>

</html>`

func NewHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:] //will get all elements of path from the first index onwards. since go is 0-indexed, this will trim the /

	//the ,ok usage here says 'only do this if you find something in the map'
	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)                                                    //log actual error message for developer
			http.Error(w, "Something went wrong...", http.StatusInternalServerError) //display a standard "something went wrong" error to the user
		}
		return
	}
	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
