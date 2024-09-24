package cyoa

import (
	"cyoa/templates"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

/* TYPES */
type ChapterOption struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type Chapter struct {
	Title      string          `json:"title"`
	Paragraphs []string        `json:"story"`
	Options    []ChapterOption `json:"options"`
}

type Story map[string]Chapter

type handler struct {
	s      Story
	t      *template.Template
	pathFn func(r *http.Request) string
}

type HandlerOption func(h *handler)

/* UTILS */
func NewStory(r io.Reader) (*Story, error) {
	d := json.NewDecoder(r)
	var story Story

	if err := d.Decode(&story); err != nil {
		return nil, fmt.Errorf("failed to parse specified json file, error: %s", err)
	}

	return &story, nil
}

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathFunc(pathFn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathFn = pathFn
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s: s, t: templates.DefaultTemplate, pathFn: defaultPathFn}

	for _, opt := range opts {
		opt(&h)
	}

	return h
}

func defaultPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	return path
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	path := h.pathFn(r)

	if chapter, ok := h.s[path]; ok {
		err = h.t.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
