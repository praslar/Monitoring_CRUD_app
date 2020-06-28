package server

import (
	"log"
	"net/http"
)

func ListenAndServe(router http.Handler) {
	address := ":8080"
	srv := &http.Server{
		Addr:    address,
		Handler: router,
	}
	log.Println("Listening on: ", address)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic("Err listenning on: ", err)
	}
}
