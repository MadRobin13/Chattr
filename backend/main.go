package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool {return true},
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println(err)
			return
		}

	}
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}

	reader(ws)
}

func setUpRoutes() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "Simple Server")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Errorf("Error reading body: %v", err)
		}
		defer r.Body.Close()

		fmt.Fprintf(w, "message: %v\n", string(body))
	})
	
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat App v0.01")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}