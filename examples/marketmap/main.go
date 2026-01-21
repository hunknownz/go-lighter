package main

import (
	"context"
	"fmt"

	"github.com/hunknownz/go-lighter/lighter"
)

func main() {
	markets, err := lighter.FetchMarkets(context.Background(), "https://explorer.elliot.ai")
	if err != nil {
		panic(err)
	}

	for _, m := range markets {
		fmt.Printf("%s -> %d\n", m.Symbol, m.MarketIndex)
	}
}
