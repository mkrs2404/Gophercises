package adventure

import (
	"html/template"
	"net/http"
	"strings"
)

var tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))

var defaultHandlerTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Choose Your Own Adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
            <li><a href="/{{.Chapter}}">{{.Text}}</a> </li>
        {{end}}
    </ul>
</body>
</html>`

type OptionalFunc func(h *handler)

func NewHandler(s Story, options ...OptionalFunc) http.Handler {
	h := handler{s, tpl}
	for _, opt := range options {
		opt(&h)
	}
	return h
}

type handler struct {
	s Story
	t *template.Template
}

func WithTemplate(t *template.Template) OptionalFunc {
	return func(h *handler) {
		h.t = t
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimSpace(r.URL.Path)
	path = path[1:]
	var err error
	if chapter, ok := h.s[path]; ok {
		err = h.t.Execute(w, chapter)
	} else {
		err = h.t.Execute(w, h.s["intro"])
	}
	if err != nil {
		http.Error(w, "Something went wrong...", http.StatusInternalServerError)
	}
}

type Story map[string]Chapter

type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}
