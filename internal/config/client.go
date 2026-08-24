package config

import (
	"context"
	"fmt"

	"github.com/juli3nk/simplelogin-cli/pkg/simplelogin"
)

// ClientFactory creates a SimpleLogin client from the stored configuration and
// API key. It is the default production implementation used by CLI commands.
func ClientFactory() (simplelogin.APIClient, error) {
	return ClientFactoryWithContext(context.Background())
}

// ClientFactoryWithContext creates a SimpleLogin client from the stored
// configuration and API key, using the provided context for requests.
func ClientFactoryWithContext(ctx context.Context) (simplelogin.APIClient, error) {
	cfg, err := Load()
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	apiKey, err := LoadAPIKey()
	if err != nil {
		return nil, fmt.Errorf("failed to load API key: %w", err)
	}

	client, err := simplelogin.NewClient(cfg.APIURL, apiKey)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	return client.WithContext(ctx), nil
}
