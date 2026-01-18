package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/joho/godotenv"

	"go-lighter/lighter"
)

func main() {
	loadEnv()

	baseURL := envOrDefault("LIGHTER_EXPLORER_URL", "https://explorer.elliot.ai")
	symbol := flag.String("symbol", "", "filter by symbol (e.g. BTC)")
	flag.Parse()

	markets, err := lighter.FetchMarkets(context.Background(), baseURL)
	if err != nil {
		fmt.Printf("fetch markets failed: %v\n", err)
		os.Exit(1)
	}

	if *symbol != "" {
		for _, m := range markets {
			if strings.EqualFold(m.Symbol, *symbol) {
				printEntry(m)
				return
			}
		}
		fmt.Printf("symbol not found: %s\n", *symbol)
		os.Exit(1)
	}

	sort.Slice(markets, func(i, j int) bool {
		return markets[i].MarketIndex < markets[j].MarketIndex
	})

	for _, m := range markets {
		printEntry(m)
	}
}

func printEntry(m lighter.MarketEntry) {
	fmt.Printf("symbol=%s market_index=%d\n", m.Symbol, m.MarketIndex)
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
