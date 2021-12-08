package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
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
        {{.}}
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
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		panic(err)
	}
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
