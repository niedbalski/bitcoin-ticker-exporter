package exchanges

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"
	"strings"
)

type BudaExchange struct {
	BaseExchange
}


func NewBudaExchange(config ExchangeConfig) (*BudaExchange, error) {
	var exchange BudaExchange
	exchange.Config = config
	exchange.Namespace = "buda"
	exchange.MetricsByMarket = []string{"ask", "bid", "lasttrade", "volume", "trades", "low", "high", "opening"}
	return &exchange, nil
}

type TickerBody struct {
	MarketId string `json:"market_id"`
	LastPrice []string `json:"last_price"`
	MinAsk []string `json:"min_ask"`
	MaxBid []string `json:"max_bid"`
	Volume []string `json:"volume"`
	PriceVariation24h string `json:"price_variation_24h"`
	PriceVariation7d string `json:"price_variation_7d"`
}

type RawTicker struct {
	Body TickerBody `json:"ticker"`
}

func (exporter *BudaExchange) GetTicker(market Market) (*Ticker, error) {
	var rawTicker RawTicker

	resp, err := http.Get(fmt.Sprintf("https://www.buda.com/api/v2/markets/%s/ticker.json", market.Code))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &rawTicker)
	if err != nil {
		return nil, err
	}

	ask, err := strconv.ParseFloat(strings.TrimSuffix(rawTicker.Body.MinAsk[0], ".0"), 64)
	if err != nil {
		return nil, err
	}


	bid, err := strconv.ParseFloat(strings.Replace(rawTicker.Body.MaxBid[0], ".0", "", -1), 64)
	if err != nil {
		return nil, err
	}


	lastTrade, err := strconv.ParseFloat(strings.TrimRight(rawTicker.Body.LastPrice[0], ".0"), 64)
	if err != nil {
		return nil, err
	}


	volume, err := strconv.ParseFloat(rawTicker.Body.Volume[0], 64)
	if err != nil {
		return nil, err
	}

	fmt.Printf("%f", ask)

	return &Ticker{ Ask: ask, Bid: bid, LastTrade: lastTrade, Volume: volume, Trades: volume, Low: float64(ask), High: bid, Opening: 0.0 }, nil
}