package monitor

import (
	"github.com/gorilla/websocket"
)

type (
	Client struct {
		// socket is the web socket for this client.
		socket *websocket.Conn
		// send is a channel on which messages are sent.
		send chan ChangeInfo
	}
)

func (c *Client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}

func (c *Client) read() {
	defer c.socket.Close()
	for {
		_, _, err := c.socket.ReadMessage()
		if err != nil {
			return
		}
	}
}
