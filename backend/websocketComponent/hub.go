package websocketComponent

import "fmt"

var Hubs = Hub{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
}

type Hub struct {
	clients map[*Client]bool

	broadcast chan []byte

	register chan *Client

	unregister chan *Client
}

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			h.clients[conn] = true
			fmt.Println("First -> Assign Client to Hub")
		case conn := <-h.unregister:
			if _, existed := h.clients[conn]; existed {
				close(conn.send)
				delete(h.clients, conn)
			}
		case message := <-h.broadcast:
			var index int
			for conn := range h.clients {
				if _, existed := h.clients[conn]; existed {
					index += 1
					conn.send <- message
					fmt.Println("Third -> Broadcast to User", index)
				}
			}
		}
	}
}

func (h *Hub) Registry(client *Client) {
	h.register <- client
}

func (h *Hub) send(message []byte, sender *Client) {
	for conn := range h.clients {
		if conn != sender {
			conn.send <- message
		}
	}
}
