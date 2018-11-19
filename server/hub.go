package main

import (
	"fmt"

	"github.com/dgrr/fastws"
)

// A Hub is a central point for WebSocket communication
type Hub struct {
	clients map[*Client]bool

	// Message ingress
	broadcast chan []byte

	// Register a client to this hub
	register chan *Client

	// Unregister a client from this hub
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (hub *Hub) run() {
	for {
		select {
		case client := <-hub.register:
			fmt.Printf("Registering client %v\n", &client)
			hub.clients[client] = true
		case client := <-hub.unregister:
			fmt.Printf("Unregistering client %v\n", &client)
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}
		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}

// HandleWebSocket processes a WebSocket connection
//
// This involves creating a client, registering that client to
// the primary hub, and then starting the write/read cycle.
func HandleWebSocket(socket *fastws.Conn, hub *Hub) {
	fmt.Printf("Received request to register\n")

	client := &Client{hub: hub, connection: socket, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.write()
	client.read()
}
