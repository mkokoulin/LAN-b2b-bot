package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2/google"

	"lan_b2b_bot/internal/config"
)

func NewGoogleClient(ctx context.Context, gcc config.GoogleCloudConfig, scope ...string) (*http.Client, error) {
	byteValue, err := json.Marshal(gcc)
	if err != nil {
		fmt.Println(err)
	}

	config, err := google.JWTConfigFromJSON(byteValue, scope...)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	// create client with config and context
	return config.Client(ctx), nil
}