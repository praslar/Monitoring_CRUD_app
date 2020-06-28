package template

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type (
	Handler struct {
		once     sync.Once
		filename string
		templ    *template.Template
	}
)

func NewHandler(filename string) *Handler {
	return &Handler{
		filename: filename,
	}
}

func (h *Handler) ServeTemplate(w http.ResponseWriter, r *http.Request) {

	h.once.Do(func() {
		h.templ = template.Must(template.ParseFiles(filepath.Join("template", h.filename)))
	})
	h.templ.Execute(w, nil)
}
