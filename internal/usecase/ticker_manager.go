package usecase

import (
	"log"
	"strings"
	"time"

	"github.com/islu/crypto-bar/internal/adapter/cex"
)

type TickerManager struct {
	binanceCaller   cex.BinanceCaller
	CurrIntervalIdx int
	CurrSymbolIdx   int
	Intervals       []time.Duration
	Symbols         []string
}

func NewTickerManager() TickerManager {
	return TickerManager{
		binanceCaller:   cex.NewBinanceCaller(),
		CurrSymbolIdx:   0,
		CurrIntervalIdx: 0,
		Intervals:       []time.Duration{30 * time.Second, 1 * time.Minute, 5 * time.Minute, 15 * time.Minute, 30 * time.Minute},
		Symbols:         []string{"BTC", "ETH", "BNB", "SOL", "SUI"},
	}
}

func (m *TickerManager) GetCurrSymbol() string {
	return m.Symbols[m.CurrSymbolIdx]
}

func (m *TickerManager) GetCurrInterval() time.Duration {
	return m.Intervals[m.CurrIntervalIdx]
}

func (m *TickerManager) GetTickerPrice(symbol string) string {
	tickerPrice, err := m.binanceCaller.GetTicketPrice(symbol)

	if err != nil {
		log.Println(err)
		return "-1"
	}

	// Trim trailing zeros
	price := strings.TrimRight(strings.TrimRight(tickerPrice.Price, "0"), ".")

	return price
}
