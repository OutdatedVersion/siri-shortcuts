package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/dgrr/fastws"
)

const (
	pingInterval = 5 * time.Second
)

type Client struct {
	hub *Hub

	connection *fastws.Conn

	send chan []byte
}

func (client *Client) read() {
	defer func() {
		client.hub.unregister <- client
		client.connection.Close("Server shutting down")
	}()

	var message []byte
	var err error

	for {
		_, message, err = client.connection.ReadMessage(message)

		if err != nil {
			fmt.Printf("Socket closed %v\n", err)

			break
		}

		message = bytes.TrimSpace(message)
		client.hub.broadcast <- message
	}
}

func (client *Client) write() {
	ticker := time.NewTicker(pingInterval)

	defer func() {
		ticker.Stop()
		client.connection.Close("Server closing")
	}()

	for {
		select {
		case message, ok := <-client.send:
			if !ok {
				client.connection.WriteString("The hub closed your channel")
				return
			}

			client.connection.Write(message)

			n := len(client.send)
			for i := 0; i < n; i++ {
				client.connection.Write(<-client.send)
			}
		case <-ticker.C:
			client.connection.WriteString("PING")
		}
	}
}
