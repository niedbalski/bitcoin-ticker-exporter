package exchanges

import (
	"github.com/beldur/kraken-go-api-client"
	"strconv"
)

type KrakenExchange struct {
	BaseExchange
	APIClient *krakenapi.KrakenApi
}

func NewKrakenExchange(config ExchangeConfig) (*KrakenExchange, error) {
	var exchange KrakenExchange
	exchange.APIClient = krakenapi.New(config.API["key"], config.API["secret"])
	exchange.Config = config
	exchange.Namespace = "kraken"
	exchange.MetricsByMarket = []string{"ask", "bid", "lasttrade", "volume", "trades", "low", "high", "opening"}
	return &exchange, nil
}

func (exchange *KrakenExchange) GetTicker(market Market) (*Ticker, error) {
	var ticker map[string]interface{}
	result, err := exchange.APIClient.Query("Ticker", map[string]string{
		"pair": market.Code,
	})

	if err != nil {
		return nil, err
	}

	ticker = result.(map[string]interface{})[krakenapi.XXBTZUSD].(map[string]interface{})

	ask, err := strconv.ParseFloat(ticker["a"].([]interface{})[0].(string), 64)
	if err != nil {
		return nil, err
	}

	bid, err := strconv.ParseFloat(ticker["b"].([]interface{})[0].(string), 64)
	if err != nil {
		return nil, err
	}

	lastTrade, err := strconv.ParseFloat(ticker["l"].([]interface{})[0].(string), 64)
	if err != nil {
		return nil, err
	}

	volume, err := strconv.ParseFloat(ticker["v"].([]interface{})[0].(string), 64)
	if err != nil {
		return nil, err
	}

	trades := ticker["t"].([]interface{})[0].(float64)

	low, err := strconv.ParseFloat(ticker["l"].([]interface{})[0].(string), 64)
	if err != nil {
		return nil, err
	}

	high, err := strconv.ParseFloat(ticker["h"].([]interface{})[0].(string), 64)
	if err != nil {
		return nil, err
	}

	opening, err := strconv.ParseFloat(ticker["o"].(interface{}).(string), 64)
	if err != nil {
		return nil, err
	}

	return &Ticker{ Ask: ask, Bid: bid, LastTrade: lastTrade, Volume: volume, Trades: trades, Low: low, High: high, Opening: opening }, nil
}
