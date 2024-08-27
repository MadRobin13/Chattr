package main

import (
	"fmt"
	"net/http"

	"github/MadRobin13/Go-and-React-Website/pkg/websocket"
)





func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setUpRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chattr v0.01")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}