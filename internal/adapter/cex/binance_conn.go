package cex

import (
	"context"

	binance_connector "github.com/binance/binance-connector-go"
	"github.com/islu/crypto-bar/internal/domain/models"
)

type BinanceCaller struct {
	apiKey    string
	secretKey string
	baseURL   string
}

func NewBinanceCaller() BinanceCaller {

	return BinanceCaller{
		apiKey:    "",
		secretKey: "",
		baseURL:   "https://api.binance.com",
	}
}

func (c *BinanceCaller) GetTicketPrice(symbol string) (*models.TickerPrice, error) {

	client := binance_connector.NewClient(c.apiKey, c.secretKey, c.baseURL)

	pair := symbol + "USDT"

	// TickerPrice
	tickerPrice, err := client.NewTickerPriceService().
		Symbol(pair).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	result := &models.TickerPrice{
		Symbol:      symbol,
		TradingPair: tickerPrice[0].Symbol,
		Price:       tickerPrice[0].Price,
	}
	return result, nil
}
