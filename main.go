package main

import (
	"log"
	"net/http"
	"sendToTelegram/api"
)

func main() {
	http.HandleFunc("/send", api.HandleTelegramSend)
	port := "3000"

	log.Printf("Server starting on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
