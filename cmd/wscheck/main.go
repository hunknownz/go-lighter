package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

type orderBookUpdate struct {
	Type string `json:"type"`
}

func main() {
	loadEnv()

	wsURL := envOrDefault("LIGHTER_WS_URL", "wss://testnet.zklighter.elliot.ai/stream")
	market := flag.Int("market", 0, "market index to subscribe")
	timeout := flag.Duration("timeout", 5*time.Second, "time to wait for a response")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, wsURL, nil)
	if err != nil {
		fmt.Printf("dial websocket failed: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_ = conn.Close()
	}()

	sub := map[string]string{
		"type":    "subscribe",
		"channel": fmt.Sprintf("order_book/%d", *market),
	}
	if err := conn.WriteJSON(sub); err != nil {
		fmt.Printf("subscribe failed: %v\n", err)
		os.Exit(1)
	}

	_ = conn.SetReadDeadline(time.Now().Add(*timeout))
	_, msg, err := conn.ReadMessage()
	if err != nil {
		fmt.Printf("read failed: %v\n", err)
		os.Exit(1)
	}

	var update orderBookUpdate
	if err := json.Unmarshal(msg, &update); err == nil && update.Type != "" {
		fmt.Printf("ws ok: %s\n", update.Type)
		return
	}

	fmt.Printf("ws ok: %s\n", string(msg))
}

func loadEnv() {
	_ = godotenv.Load()
}

func envOrDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}
