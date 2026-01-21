package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/hunknownz/go-lighter/lighter"
)

func main() {
	baseURL := envOrDefault("LIGHTER_BASE_URL", "https://testnet.zklighter.elliot.ai")
	chainID := envUint32("LIGHTER_CHAIN_ID", 1)
	accountIndex := envInt64("LIGHTER_ACCOUNT_INDEX", 0)
	apiKeyIndex := envUint8("LIGHTER_API_KEY_INDEX", 0)
	apiPrivateKey := os.Getenv("LIGHTER_API_PRIVATE_KEY")
	if apiPrivateKey == "" {
		fmt.Println("LIGHTER_API_PRIVATE_KEY is required")
		os.Exit(1)
	}

	signer, err := lighter.NewSigner(lighter.SignerConfig{
		BaseURL:       baseURL,
		ChainID:       chainID,
		AccountIndex:  accountIndex,
		APIKeyIndex:   apiKeyIndex,
		APIPrivateKey: apiPrivateKey,
	})
	if err != nil {
		fmt.Printf("init signer failed: %v\n", err)
		os.Exit(1)
	}

	token, err := signer.AuthToken(time.Now().Add(10 * time.Minute))
	if err != nil {
		fmt.Printf("create auth token failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(token)
}

func envOrDefault(key, def string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return def
}

func envInt64(key string, def int64) int64 {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	parsed, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return def
	}
	return parsed
}

func envUint32(key string, def uint32) uint32 {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	parsed, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return def
	}
	return uint32(parsed)
}

func envUint8(key string, def uint8) uint8 {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	parsed, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return def
	}
	return uint8(parsed)
}
