package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// ProcessAction makes an attempt to run the provided action
func ProcessAction(context *fasthttp.RequestCtx, hub *Hub) {
	var action = context.UserValue("action").(string)

	fmt.Fprintf(
		context,
		"Received request to perform action %s",
		context.UserValue("action"))

	if !IsAuthorized(context, "computers."+action) {
		return
	}

	hub.broadcast <- []byte(action)
}
