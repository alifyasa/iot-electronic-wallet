package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Not safe for production; accept connections from any origin
	},
}

// A thread-safe map to hold WSConnections
var WSConnections = struct {
	sync.RWMutex
	items map[*websocket.Conn]bool
}{items: make(map[*websocket.Conn]bool)}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	// Register the new connection
	WSConnections.Lock()
	WSConnections.items[conn] = true
	WSConnections.Unlock()

	// Test Reply
	err = conn.WriteMessage(websocket.TextMessage, []byte("Connected to Server"))
	if err != nil {
		log.Println(err)
		return
	}

	// Wait for the connection to close
	for {
		if _, _, err := conn.NextReader(); err != nil {
			WSConnections.Lock()
			log.Printf("Disconnecting %s\n", conn.RemoteAddr())
			delete(WSConnections.items, conn)
			WSConnections.Unlock()
			break
		}
	}
}

/* Broadcast message to every WS connection */
func SendMessageToWebSockets(message string) {
	WSConnections.RLock()
	defer WSConnections.RUnlock()
	for conn := range WSConnections.items {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
			log.Println(err)
			// Remove the connection if it's not valid anymore
			WSConnections.Lock()
			delete(WSConnections.items, conn)
			WSConnections.Unlock()
		}
	}
}
