package monitor

import "github.com/gorilla/websocket"

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize,
	WriteBufferSize: socketBufferSize}

type (
	Room struct {
		forward chan ChangeInfo
		join    chan *Client
		leave   chan *Client
		clients map[*Client]bool
	}
)

func NewRoom() *Room {
	return &Room{
		forward: make(chan ChangeInfo),
		join:    make(chan *Client),
		leave:   make(chan *Client),
		clients: make(map[*Client]bool),
	}
}

func (r *Room) Join(client *Client) {
	r.join <- client
}

func (r *Room) Leave(client *Client) {
	r.leave <- client
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.join:
			// joining
			r.clients[client] = true
		case client := <-r.leave:
			// leaving
			delete(r.clients, client)
			close(client.send)
		case msg := <-r.forward:
			// forward message to all clients
			for client := range r.clients {
				client.send <- msg
			}
		}
	}
}

func (r *Room) Write(msg ChangeInfo) {
	r.forward <- msg
}
