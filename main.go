package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func websocketHandler(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	for {

		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Received message: %s", message)

		err = conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	r := gin.Default()
	r.GET("/websocket", websocketHandler)
	r.Run("localhost:3000")
}
