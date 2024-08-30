package websocket

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int `json:"type"`
	Body string `json:"body"`
}

func (c *Client) Read() { 
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message
		fmt.Printf("Message recieved %+v\n", message)
	}
}

func (pool *Pool) Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	for {
		select {
		case client := <- pool.Register:
			pool.Clients[client] = true
			fmt.Println("size of connecton pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: "New user joined..."})
			} 

		case client := <- pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(Message{Type: 1, Body: "User disconnected..."}); err != nil {
					fmt.Println(err)
					return
				}
			}

		case message := <- pool.Broadcast:
			fmt.Println("Sending message to all pool clients")
			for client, _ := range pool.Clients{
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("URI")))
			if err != nil {
				log.Fatal("couldn't apply uri: %v", err)
			}
			
			client.Database("Chattr").Collection("Messages").InsertOne(ctx, message)
		}
	}
}