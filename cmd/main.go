package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Serve static files from the `static` directory
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("static"))

	mux.Handle("/", fs)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/ws", WsHandler)

	loggedMux := LoggingMiddleware(mux)

	fmt.Println("Server starting on port 8080...")
	log.Fatal(
		http.ListenAndServe(":8080", loggedMux),
	)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("(%s) Hello\n", r.RemoteAddr)
	msg := fmt.Sprintf("[%s] Hello!\n", time.Now())
	SendMessageToWebSockets(msg)
	fmt.Fprintf(w, "Message sent to WebSocket clients")
}
