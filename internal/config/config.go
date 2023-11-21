package config

import (
	"fmt"
	"log"

	"github.com/caarlos0/env/v6"
)

type GoogleConfig struct {
	Scope string `env:"SCOPE" json:"SCOPE"`
}

type BuilderTable struct {
	SpreadsheetId string `env:"BUILDER_SPREADSHEET_ID" json:"BUILDER_SPREADSHEET_ID"`
	ReadRange string `env:"BUILDER_READ_RANGE" json:"BUILDER_READ_RANGE"`
}

type RequestsTable struct {
	SpreadsheetId string `env:"REQUESTS_SPREADSHEET_ID" json:"REQUESTS_SPREADSHEET_ID"`
	ReadRange string `env:"REQUESTS_READ_RANGE" json:"REQUESTS_READ_RANGE"`
}

type GoogleSheetsConfig struct {
	Builder BuilderTable
	Requests RequestsTable
}

type Config struct {
	TelegramToken string `env:"TELEGRAM_TOKEN" json:"TELEGRAM_TOKEN"`
	Google GoogleConfig
	GoogleSheets GoogleSheetsConfig
}

func New() (*Config, error) {
	cfg := Config{}
	
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	gc := GoogleConfig{}

	err = env.Parse(&gc)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	gsc := GoogleSheetsConfig{}

	builderTable := BuilderTable{}
	requestsTable := RequestsTable{}

	err = env.Parse(&builderTable)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	err = env.Parse(&requestsTable)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	gsc.Builder = builderTable
	gsc.Requests = requestsTable
	
	cfg.Google = gc
	cfg.GoogleSheets = gsc

	log.Default().Printf("TELEGRAM_TOKEN: %v", cfg.TelegramToken)

	log.Default().Printf("SCOPE: %v", cfg.Google.Scope)

	log.Default().Printf("BUILDER_SPREADSHEET_ID: %v", cfg.GoogleSheets.Builder.SpreadsheetId)
	log.Default().Printf("BUILDER_READ_RANGE: %v", cfg.GoogleSheets.Builder.ReadRange)
	log.Default().Printf("REQUESTS_SPREADSHEET_ID: %v", cfg.GoogleSheets.Requests.SpreadsheetId)
	log.Default().Printf("REQUESTS_READ_RANGE: %v", cfg.GoogleSheets.Requests.ReadRange)

	return &cfg, nil
}