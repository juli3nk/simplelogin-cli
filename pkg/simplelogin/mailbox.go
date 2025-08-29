package simplelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Mailbox struct {
	Email             string `json:"email"`
	ID                int    `json:"id"`
	Default           bool   `json:"default"`
	CreationTimestamp int    `json:"creation_timestamp"`
	NBAlias           int    `json:"nb_alias"`
	Verified          bool   `json:"verified"`
}

// MailboxResponse represents the response for listing mailboxes
type MailboxResponse struct {
	Mailboxes []Mailbox `json:"mailboxes"`
}

type MailboxCreateResponse struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Verified bool   `json:"verified"`
	Default  bool   `json:"default"`
}

type MailboxDeleteOptions struct {
	TransferAliasesTo *int `json:"transfer_aliases_to"`
}

func (c *Client) GetMailboxes() ([]Mailbox, error) {
	endpoint := "/v2/mailboxes"

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result MailboxResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Mailboxes, nil
}

func (c *Client) CreateMailbox(email string) (*MailboxCreateResponse, error) {
	// Validate required fields
	if email == "" {
		return nil, &ValidationError{Field: "email", Message: "email is required"}
	}

	endpoint := "/mailboxes"

	jsonData, err := json.Marshal(map[string]string{"email": email})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Mailbox data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result MailboxCreateResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteMailbox(mailboxID int, options MailboxDeleteOptions) error {
	// Validate mailbox ID
	if mailboxID <= 0 {
		return &ValidationError{Field: "mailboxID", Message: "mailbox ID must be positive"}
	}

	endpoint := fmt.Sprintf("/mailboxes/%d", mailboxID)

	jsonData, err := json.Marshal(options)
	if err != nil {
		return fmt.Errorf("failed to marshal MailboxDeleteOptions data: %w", err)
	}

	resp, err := c.doRequest(http.MethodDelete, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	return c.handleResponse(resp, nil)
}
