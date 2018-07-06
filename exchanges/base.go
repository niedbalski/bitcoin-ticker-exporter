package exchanges

import (
	"github.com/prometheus/client_golang/prometheus"
	"reflect"
	"strings"
	"fmt"
)

type Market struct {
	Name string `yaml:"name"`
	Code string `yaml:"code"`
}

type Ticker struct {
	Ask float64
	Bid float64
	LastTrade float64
	Volume float64
	Trades float64
	Low float64
	High float64
	Opening float64
}

func (ticker *Ticker) GetAttribute(name string) (float64) {
	t := reflect.Indirect(reflect.ValueOf(&ticker)).Elem()
	typeOfT := t.Type()
	for i := 0; i < t.NumField(); i++ {
		if strings.ToLower(typeOfT.Field(i).Name) == name {
			return t.Field(i).Float()
		}
	}
	return 0.0
}

type ExchangeConfig struct {
	Name   string `yaml:"name"`
	API map[string]string `yaml:"api"`
	Markets []Market `yaml:"markets"`
}

type Exchange interface {
	GetTicker(market Market) (*Ticker, error)
	GetName() (string)
	GetMarkets() ([]Market)
	GetMetricsByMarket() ([]string)
	GetMetrics() (map[string]*prometheus.Desc)
 	RegisterMetrics()
}

type BaseExchange struct {
	Config ExchangeConfig
	Metrics map[string]*prometheus.Desc
	Namespace string
	MetricsByMarket []string
}

func (exchange *BaseExchange) GetMarkets() ([]Market) {
	return exchange.Config.Markets
}

func (exchange *BaseExchange) GetMetricsByMarket() ([]string) {
	return exchange.MetricsByMarket
}

func (exchange *BaseExchange) GetName() (string) {
	return exchange.Config.Name
}

func (exchange *BaseExchange) GetMetrics() (map[string]*prometheus.Desc) {
	return exchange.Metrics
}

func (exchange *BaseExchange) GetNamespace() (string) {
	return exchange.Namespace
}

func (exchange *BaseExchange) RegisterMetrics() {
	exchange.Metrics = map[string]*prometheus.Desc{}

	for _, market := range exchange.GetMarkets() {
		for _, metric := range exchange.GetMetricsByMarket() {
			metricKey := fmt.Sprintf("%s_%s", strings.ToLower(market.Name), metric)
			if _, ok := exchange.Metrics[metricKey]; !ok {
				exchange.Metrics[metricKey] = prometheus.NewDesc(
					prometheus.BuildFQName(exchange.GetNamespace(), "", metricKey),
					metricKey, nil, nil)
			}
		}
	}
}

