package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/zalando/go-keyring"
)

const (
	service = "simplelogin-cli"
	user    = "default"
)

func SaveAPIKey(apiKey string) error {
	err := keyring.Set(service, user, apiKey)
	if err == nil {
		return nil
	}

	// Fallback: store the API key in a local file with restricted permissions.
	// The keyring is preferred, but some environments (CI, containers, SSH
	// sessions) do not provide one. The file is created with 0600 permissions.
	fmt.Println("⚠️  Warning: keyring not available, falling back to local file storage.")
	fmt.Println("   The API key will be stored in ~/.config/simplelogin-cli/credentials.json")

	return saveApiKeyFile(apiKey)
}

func LoadAPIKey() (string, error) {
	apiKey, err := keyring.Get(service, user)
	if err == nil {
		return apiKey, nil
	}

	// Fallback
	return loadApiKeyFile()
}

func DeleteAPIKey() error {
	err := keyring.Delete(service, user)
	if err == nil {
		return nil
	}

	// Fallback
	return deleteApiKeyFile()
}

//
// --- Fallback : fichier local ---
//

func credsPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".config", "simplelogin-cli", "credentials.json"), nil
}

func saveApiKeyFile(apiKey string) error {
	path, err := credsPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}

	data, err := json.MarshalIndent(map[string]string{"api_key": apiKey}, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

func loadApiKeyFile() (string, error) {
	path, err := credsPath()
	if err != nil {
		return "", err
	}

	//nolint:gosec
	// false positive: reading config file from trusted user path
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("no API key found")
	} else if err != nil {
		return "", err
	}

	var obj map[string]string
	if err := json.Unmarshal(data, &obj); err != nil {
		return "", err
	}

	return obj["api_key"], nil
}

func deleteApiKeyFile() error {
	path, err := credsPath()
	if err != nil {
		return err
	}
	return os.Remove(path)
}
