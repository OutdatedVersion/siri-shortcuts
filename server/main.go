package main

import (
	"flag"
	"log"

	"github.com/dgrr/fastws"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

var (
	address = flag.String("a", ":8080", "TCP address to listen to")
)

func createWebSocketWrapper(hub *Hub) fastws.RequestHandler {
	return fastws.RequestHandler(func(conn *fastws.Conn) {
		HandleWebSocket(conn, hub)
	})
}

func createHTTPWrapper(hub *Hub) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(context *fasthttp.RequestCtx) {
		ProcessAction(context, hub)
	})
}

func main() {
	flag.Parse()

	log.Printf("Starting gateway (target address %s)", *address)

	router := router.New()
	hub := newHub()

	go hub.run()

	router.GET("/api/hub", fastws.Upgrade(createWebSocketWrapper(hub)))
	router.POST("/api/:action", createHTTPWrapper(hub))

	if err := fasthttp.ListenAndServe(*address, router.Handler); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}
