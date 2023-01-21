package websocketComponent

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub *Hub

	socket *websocket.Conn

	send chan []byte
}

func ConstructClient(hub *Hub, socket *websocket.Conn, send chan []byte) *Client {
	return &Client{
		hub:    hub,
		socket: socket,
		send:   send,
	}
}

func (c *Client) Read() {
	defer func() {
		c.hub.unregister <- c
		c.socket.Close()
	}()
	for {
		_, message, err := c.socket.ReadMessage()

		if err != nil {
			if websocket.IsCloseError(err, 1001) {
				fmt.Println("User went offline")
			}
			c.hub.unregister <- c
			c.socket.Close()
			break
		}
		fmt.Println("Second -> Send message from client to Hub")
		c.hub.broadcast <- message
	}
}

func (c *Client) Write() {
	defer func() {
		c.socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
			fmt.Println("Forth -> Submit message from channel")
		}
	}
}
