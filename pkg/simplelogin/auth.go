package simplelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// LoginRequest represents the login request payload
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Device   string `json:"device"`
}

// LoginResponse represents the login response
type LoginResponse struct {
	APIKey string `json:"api_key"`
}

// Login authenticates a user and returns an API key
func Login(email, password, device string) (string, error) {
	client := &http.Client{}

	data := LoginRequest{
		Email:    email,
		Password: password,
		Device:   device,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("failed to marshal login data: %w", err)
	}

	resp, err := client.Post(BaseURL+"/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("login request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := json.Marshal(resp.Body)
		return "", fmt.Errorf("login failed (status %d): %s", resp.StatusCode, string(body))
	}

	var result LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode login response: %w", err)
	}

	return result.APIKey, nil
}
