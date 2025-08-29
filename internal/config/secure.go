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

func SaveApiKey(apiKey string) error {
	err := keyring.Set(service, user, apiKey)
	if err == nil {
		return nil
	}

	// Fallback : fichier ~/.app-cli/credentials.json
	fmt.Println("⚠️  Warning: keyring not available, falling back to local file storage.")

	return saveApiKeyFile(apiKey)
}

func LoadApiKey() (string, error) {
	apiKey, err := keyring.Get(service, user)
	if err == nil {
		return apiKey, nil
	}

	// Fallback
	return loadApiKeyFile()
}

func DeleteApiKey() error {
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

	data, _ := json.MarshalIndent(map[string]string{"api_key": apiKey}, "", "  ")
	return os.WriteFile(path, data, 0600)
}

func loadApiKeyFile() (string, error) {
	path, err := credsPath()
	if err != nil {
		return "", err
	}

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
