package lighter

import (
	"encoding/hex"
	"fmt"
	"time"

	coreclient "github.com/elliottech/lighter-go/client"
	corehttp "github.com/elliottech/lighter-go/client/http"
	"github.com/elliottech/lighter-go/types"
)

type SignerConfig struct {
	BaseURL       string
	ChainID       uint32
	AccountIndex  int64
	APIKeyIndex   uint8
	APIPrivateKey string
}

type Signer struct {
	txClient *coreclient.TxClient
}

type CreateOrderRequest struct {
	MarketIndex      int16
	ClientOrderIndex int64
	BaseAmount       int64
	Price            uint32
	IsAsk            bool
	Type             uint8
	TimeInForce      uint8
	ReduceOnly       bool
	TriggerPrice     uint32
	OrderExpiry      int64
}

type SignedTx struct {
	TxType uint8
	TxInfo string
	TxHash string
	SigHex string
}

func NewSigner(cfg SignerConfig) (*Signer, error) {
	if cfg.AccountIndex <= 0 {
		return nil, fmt.Errorf("account index must be > 0")
	}
	if cfg.APIPrivateKey == "" {
		return nil, fmt.Errorf("api private key is required")
	}

	httpClient := corehttp.NewClient(cfg.BaseURL)
	txClient, err := coreclient.CreateClient(
		httpClient,
		cfg.APIPrivateKey,
		cfg.ChainID,
		cfg.APIKeyIndex,
		cfg.AccountIndex,
	)
	if err != nil {
		return nil, err
	}

	return &Signer{txClient: txClient}, nil
}

func (s *Signer) AuthToken(deadline time.Time) (string, error) {
	return s.txClient.GetAuthToken(deadline)
}

func (s *Signer) SignCreateOrder(req CreateOrderRequest, opts *types.TransactOpts) (*SignedTx, error) {
	tx, err := s.txClient.GetCreateOrderTransaction(&types.CreateOrderTxReq{
		MarketIndex:      req.MarketIndex,
		ClientOrderIndex: req.ClientOrderIndex,
		BaseAmount:       req.BaseAmount,
		Price:            req.Price,
		IsAsk:            boolToUint8(req.IsAsk),
		Type:             req.Type,
		TimeInForce:      req.TimeInForce,
		ReduceOnly:       boolToUint8(req.ReduceOnly),
		TriggerPrice:     req.TriggerPrice,
		OrderExpiry:      req.OrderExpiry,
	}, opts)
	if err != nil {
		return nil, err
	}

	info, err := tx.GetTxInfo()
	if err != nil {
		return nil, err
	}

	return &SignedTx{
		TxType: tx.GetTxType(),
		TxInfo: info,
		TxHash: tx.GetTxHash(),
		SigHex: hex.EncodeToString(tx.Sig),
	}, nil
}

func boolToUint8(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}
