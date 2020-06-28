package api

import (
	"fmt"
	"net/http"

	"github.com/blogCRUDWebsocket/internal/pkg/http/router"
)

const (
	get    = http.MethodGet
	post   = http.MethodPost
	put    = http.MethodPut
	delete = http.MethodDelete
)

func NewRouter() (http.Handler, error) {
	monitorRoom := newMonitorRoom()
	monitorHandler := newMonitorHandler(monitorRoom)
	blogHandler := newBlogHandler(newBlogService(monitorRoom))
	templateHandler := newTemplateHandler("monitoring.html")

	go monitorRoom.Run()

	routes := []router.Route{}

	routes = append(routes, blogHandler.Routes()...)
	routes = append(routes, templateHandler.Routes()...)
	routes = append(routes, monitorHandler.Routes()...)

	r, err := router.New(&router.Config{
		Routes: routes,
	})

	fmt.Println("Debug: ", r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
