package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"github.com/niedbalski/bitcoin-ticker-exporter/exchanges"
)

type Market struct {
	Name string `yaml:"name"`
	Code string `yaml:"code"`
}

type API struct {
	Name string `yaml:"key"`
	Code string `yaml:"secret"`
}

type Config struct {
	Exchanges []exchanges.ExchangeConfig `yaml:"exchanges"`
}

func NewConfigFromYAML(path string) (*Config, error) {
	var config Config

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}