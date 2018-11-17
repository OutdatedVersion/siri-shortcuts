package main

import (
	"fmt"
	"os"

	"github.com/dgrr/fastws"
)

func main() {
	eventHubURI := os.Getenv("EVENT_HUB_URI")

	fmt.Printf("Using event hub URI: %s\n", eventHubURI)

	websocket, err := fastws.Dial(eventHubURI)

	if err != nil {
		fmt.Printf("Failed to open WebSocket: %s", err)
	}

	var message []byte

	for {
		_, message, err = websocket.ReadMessage(message[:0])

		if err != nil {
			if err != fastws.EOF {
				fmt.Fprintf(os.Stderr, "Error reading message: %s\n", err)
			}

			break
		}

		fmt.Printf("Received message: %s\n", message)

		switch string(message) {
		case "lock":
			LockComputer()
		}
	}
}
