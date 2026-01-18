# go-lighter

Minimal Go wrapper around the Lighter signing SDK and WebSocket order book stream.

## Auth token (local signing)

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

## Market map (Explorer REST)

Fetches `/api/markets` from the explorer service.

```bash
go run ./cmd/marketmap
go run ./cmd/marketmap -symbol BTC
```

Library usage:

```go
markets, err := lighter.FetchMarkets(context.Background(), "https://explorer.elliot.ai")
```

## WebSocket check

Subscribes to an order book channel and exits after receiving one message.

```bash
go run ./cmd/wscheck -market 0
```
