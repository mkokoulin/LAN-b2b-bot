package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type GoogleCloudConfig struct {
	Type string `env:"type" json:"type"`
	ProjectId string `env:"project_id" json:"project_id"`
	PrivateKeyId string `env:"private_key_id" json:"private_key_id"`
	PrivateKey string `env:"private_key" json:"private_key"`
	ClientEmail string `env:"client_email" json:"client_email"`
	ClientId string `env:"client_id" json:"client_id"`
	AuthUri string `env:"auth_uri" json:"auth_uri"`
	TokenUri string `env:"token_uri" json:"token_uri"`
	AuthProviderX509CertUrl string `env:"auth_provider_x509_cert_url" json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl string `env:"client_x509_cert_url" json:"client_x509_cert_url"`
}

type Config struct {
	Scope string `env:"SCOPE" json:"SCOPE"`
	TelegramToken string `env:"TELEGRAM_TOKEN" json:"TELEGRAM_TOKEN"`
	BuilderSpreadsheetId string `env:"BUILDER_SPREADSHEET_ID" json:"BUILDER_SPREADSHEET_ID"`
	BuilderReadRange string `env:"BUILDER_READ_RANGE" json:"BUILDER_READ_RANGE"`
	RequestsSpreadsheetId string `env:"REQUESTS_SPREADSHEET_ID" json:"REQUESTS_SPREADSHEET_ID"`
	RequestsReadRange string `env:"REQUESTS_READ_RANGE" json:"REQUESTS_READ_RANGE"`
	GoogleCloudConfig GoogleCloudConfig `env:"GOOGLE_CLOUD_CONFIG" json:"GOOGLE_CLOUD_CONFIG"`
}

func New() (*Config, error) {
	cfg := Config{}

	cfg.Scope = os.Getenv("SCOPE")
	if cfg.Scope == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "SCOPE")
	}
	log.Default().Printf("[LAN-TG-BOT] SCOPE: %v", cfg.Scope)

	cfg.TelegramToken = os.Getenv("TELEGRAM_TOKEN")
	if cfg.Scope == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "TELEGRAM_TOKEN")
	}
	log.Default().Printf("[LAN-TG-BOT] TELEGRAM_TOKEN: %v", cfg.TelegramToken)

	cfg.BuilderSpreadsheetId = os.Getenv("BUILDER_SPREADSHEET_ID")
	if cfg.Scope == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "BUILDER_SPREADSHEET_ID")
	}
	log.Default().Printf("[LAN-TG-BOT] BUILDER_SPREADSHEET_ID: %v", cfg.BuilderSpreadsheetId)

	cfg.BuilderReadRange = os.Getenv("BUILDER_READ_RANGE")
	if cfg.Scope == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "BUILDER_READ_RANGE")
	}
	log.Default().Printf("[LAN-TG-BOT] BUILDER_READ_RANGE: %v", cfg.BuilderReadRange)

	cfg.RequestsSpreadsheetId = os.Getenv("REQUESTS_SPREADSHEET_ID")
	if cfg.Scope == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "REQUESTS_SPREADSHEET_ID")
	}
	log.Default().Printf("[LAN-TG-BOT] REQUESTS_READ_RANGE: %v", cfg.RequestsSpreadsheetId)

	cfg.RequestsReadRange = os.Getenv("REQUESTS_READ_RANGE")
	if cfg.Scope == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "REQUESTS_READ_RANGE")
	}

	googleCloudConfigString := os.Getenv("GOOGLE_CLOUD_CONFIG")
	if googleCloudConfigString == "" {
		return nil, fmt.Errorf("environment variable %v is not set or empty", "GOOGLE_CLOUD_CONFIG")
	}
	var googleCloudConfig GoogleCloudConfig
	if err := json.Unmarshal([]byte(googleCloudConfigString), &googleCloudConfig); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	log.Default().Printf("[LAN-TG-BOT] REQUESTS_SPREADSHEET_ID: %v", cfg.RequestsReadRange)
	cfg.GoogleCloudConfig = googleCloudConfig;

	return &cfg, nil
}