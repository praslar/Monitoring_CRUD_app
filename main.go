package main

import (
	"log"

	"github.com/blogCRUDWebsocket/internal/app/api"
	"github.com/blogCRUDWebsocket/internal/pkg/http/server"
)

func main() {

	router, err := api.NewRouter()
	if err != nil {
		log.Panic("Cannot init router, err: ", err)
	}
	server.ListenAndServe(router)

}
