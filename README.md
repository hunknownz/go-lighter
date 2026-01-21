# go-lighter

Minimal Go wrapper around the Lighter signing SDK plus a small Explorer market map helper.

## Install

```bash
go get github.com/hunknownz/go-lighter
```

## SDK usage

### Create an auth token locally

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hunknownz/go-lighter/lighter"
)

func main() {
	signer, err := lighter.NewSigner(lighter.SignerConfig{
		BaseURL:       "https://testnet.zklighter.elliot.ai",
		ChainID:       1,
		AccountIndex:  123,
		APIKeyIndex:   3,
		APIPrivateKey: "0xYOUR_PRIVATE_KEY",
	})
	if err != nil {
		panic(err)
	}

	token, err := signer.AuthToken(time.Now().Add(10 * time.Minute))
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
	_ = context.Background()
}
```

### Fetch market map from the Explorer API

```go
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
```

See `examples/` for runnable versions.

## Commands

### Auth token (local signing)

Environment variables:

- `LIGHTER_API_PRIVATE_KEY`
- `LIGHTER_ACCOUNT_INDEX`
- `LIGHTER_API_KEY_INDEX`
- `LIGHTER_CHAIN_ID`
- `LIGHTER_BASE_URL` (default: testnet)
- `LIGHTER_AUTH_EXPIRY_SECONDS` (default: 600)

Run:

```bash
cp .env.example .env

LIGHTER_API_PRIVATE_KEY=... \
LIGHTER_ACCOUNT_INDEX=123 \
LIGHTER_API_KEY_INDEX=3 \
LIGHTER_CHAIN_ID=1 \
LIGHTER_BASE_URL=https://testnet.zklighter.elliot.ai \
LIGHTER_AUTH_EXPIRY_SECONDS=600 \
go run ./cmd/auth
```

The binaries load `.env` automatically via `godotenv`, so you can put secrets there and keep it out of git.

### Market map (Explorer REST)

Fetches `/api/markets` from the explorer service.

```bash
go run ./cmd/marketmap
go run ./cmd/marketmap -symbol BTC
```

### WebSocket check

Subscribes to an order book channel and exits after receiving one message.

```bash
go run ./cmd/wscheck -market 0
```
