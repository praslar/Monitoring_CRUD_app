package api

import "github.com/blogCRUDWebsocket/internal/app/template"

func newTemplateHandler(filename string) *template.Handler {
	return template.NewHandler(filename)
}
