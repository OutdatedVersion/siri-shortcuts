package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

// ProcessAction makes an attempt to run the provided action
func ProcessAction(context *fasthttp.RequestCtx, hub *Hub) {
	var action = context.UserValue("action").(string)
	var payload = context.PostBody()

	fmt.Fprintf(
		context,
		"Received request to perform action %s",
		action)

	if !IsAuthorized(context, "computers."+action) {
		return
	}

	message := &Message{
		Action:  action,
		Payload: string(payload),
	}

	json, err := json.Marshal(message)

	if err != nil {
		context.Error("Failed to marshal message into JSON", fasthttp.StatusInternalServerError)
	}

	hub.broadcast <- json
}
