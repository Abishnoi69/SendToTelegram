package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

var botToken = os.Getenv("TOKEN")

type TelegramRequest struct {
	ChatID  string `json:"chat_id"`
	Message string `json:"message"`
}

// Handler handles incoming HTTP requests
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Check for the correct endpoint
	if !strings.HasPrefix(r.URL.Path, "/send") {
		http.Error(w, "Invalid endpoint", http.StatusNotFound)
		return
	}

	// Ensure only POST requests are processed
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the request payload
	var reqData TelegramRequest
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Send the message to Telegram
	err = sendToTelegram(reqData.ChatID, reqData.Message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "Message sent to Telegram chat %s", reqData.ChatID)
}

func sendToTelegram(chatID, message string) error {
	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", botToken)
	payload := fmt.Sprintf(`{"chat_id":"%s", "text":"%s", "parse_mode":"HTML", "disable_web_page_preview": true}`, chatID, message)
	req, err := http.NewRequest("POST", telegramURL, strings.NewReader(payload))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request to Telegram: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error response from Telegram: %v", resp.Status)
	}

	log.Println("Message sent to Telegram")
	return nil
}
