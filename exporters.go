package main

import (
	"github.com/niedbalski/bitcoin-ticker-exporter/exchanges"
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
	"strings"
	"github.com/prometheus/common/log"
)

type Exporter struct {
	Exchanges []exchanges.Exchange
}

func (exporter *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, exchange := range exporter.Exchanges {
		for _, metric := range exchange.GetMetrics() {
			ch <- metric
		}
	}
}

func (exporter *Exporter) Collect(ch chan<- prometheus.Metric) {
	for _, exchange := range exporter.Exchanges {
		for _, market := range exchange.GetMarkets() {
			ticker, err := exchange.GetTicker(market)
			if err != nil {
				log.Error(err)
				continue
			}

			for _, metric := range exchange.GetMetricsByMarket() {
				metricKey := fmt.Sprintf("%s_%s", strings.ToLower(market.Name), metric)
				if _, ok := exchange.GetMetrics()[metricKey]; ok {
					ch <- prometheus.MustNewConstMetric(exchange.GetMetrics()[metricKey],
						prometheus.GaugeValue, ticker.GetAttribute(metric))
				}
			}
		}
	}
}

func NewExporter(config *Config) (*Exporter, error) {
	var exporter = Exporter{}

	for _, exchangeConfig := range config.Exchanges {
		switch exchangeConfig.Name {
		case "buda" : {
			instance, err := exchanges.NewBudaExchange(exchangeConfig)
			if err != nil {
				return nil, err
			}
			instance.RegisterMetrics()
			exporter.Exchanges = append(exporter.Exchanges, instance)
		}
		case "kraken": {
				instance, err := exchanges.NewKrakenExchange(exchangeConfig)
				if err != nil {
					return nil, err
				}
				instance.RegisterMetrics()
				exporter.Exchanges = append(exporter.Exchanges, instance)
			}
		}
	}

	return &exporter, nil
}