package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	log.Println("Starting server on port 8080")

	http.HandleFunc("/", index)
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8080", nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func ws(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			panic(err)
		}
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}
