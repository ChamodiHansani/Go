package main

import (
	"log"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

// Define a struct to represent a connected client
type Client struct {
	Conn *websocket.Conn
}

// Global list of clients and a mutex for thread-safe operations
var clients = make([]*Client, 0)
var clientsMutex sync.Mutex

func main() {
	app := fiber.New()

	// WebSocket endpoint
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		client := &Client{Conn: c}
		addClient(client)
		defer func() {
			removeClient(client)
			c.Close()
		}()

		for {
			// Read message from WebSocket
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)

			// Broadcast message to all connected clients
			broadcastMessage(mt, message)
		}
	}))

	// Serve the HTML client
	app.Static("/", "./public")

	log.Fatal(app.Listen(":3020"))
}

// Function to add a client to the list
func addClient(client *Client) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients = append(clients, client)
}

// Function to remove a client from the list
func removeClient(client *Client) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for i, c := range clients {
		if c == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

// Function to broadcast a message to all clients
func broadcastMessage(mt int, message []byte) {
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	for _, client := range clients {
		err := client.Conn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
		}
	}
}
