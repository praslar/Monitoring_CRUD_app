package monitor

import (
	"log"
	"net/http"
)

type (
	service interface {
		Run()
		Write(msg ChangeInfo)
		Leave(client *Client)
		Join(client *Client)
	}

	Handler struct {
		srv service
	}
)

func NewHandler(srv service) *Handler {
	return &Handler{
		srv: srv,
	}
}

func (r *Handler) ConnectMonitor(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	client := &Client{
		socket: socket,
		send:   make(chan ChangeInfo, messageBufferSize),
	}

	r.srv.Join(client)
	defer r.srv.Leave(client)
	go client.write()
	client.read()
}
