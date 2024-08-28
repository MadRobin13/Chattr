package main

import (
	"fmt"
	"net/http"

	"github/MadRobin13/Go-and-React-Website/pkg/websocket"
)


func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &websocket.Client {
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setUpRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Chattr v0.01")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}