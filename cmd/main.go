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
	mux.HandleFunc("/pay", PayHandler)
	mux.HandleFunc("/check-wallet", CheckWalletHandler)

	loggedMux := LoggingMiddleware(mux)

	fmt.Printf("IOT ELECTRONIC PAYMENT\n\n")
	fmt.Printf("13520135 - Muhammad Alif Putra Yasa\n")
	fmt.Printf("Dibuat untuk Ujian Tengah Semester IF4051 Semester 2 2023/2024\n\n")

	log.Println("Server starting on port 8080")
	log.Fatal(
		http.ListenAndServe("0.0.0.0:8080", loggedMux),
	)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf("[%s] Hello!\n", time.Now().Format("2006-01-02 15:04:05"))
	SendMessageToWebSockets(msg)
	fmt.Fprintf(w, "Message sent to WebSocket clients")
}
