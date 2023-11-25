package services

import (
	"context"
	"encoding/json"
	"fmt"
	"lan_b2b_bot/internal/config"
	"net/http"

	"golang.org/x/oauth2/google"
)

func NewGoogleClient(ctx context.Context, gcc config.GoogleCloudConfig, scope ...string) (*http.Client, error) {
	byteValue, err := json.Marshal(gcc)
	if err != nil {
		return nil, err
	}

	config, err := google.JWTConfigFromJSON(byteValue, scope...)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return config.Client(ctx), nil
}