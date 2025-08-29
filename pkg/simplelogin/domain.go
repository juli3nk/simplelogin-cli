package simplelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Domain represents a SimpleLogin domain
type Domain struct {
	CatchAll               bool      `json:"catch_all"`
	CreationDate           string    `json:"creation_date"`
	CreationTimestamp      int       `json:"creation_timestamp"`
	DomainName             string    `json:"domain_name"`
	ID                     int       `json:"id"`
	IsVerified             bool      `json:"is_verified"`
	Mailboxes              []Mailbox `json:"mailboxes"`
	Name                   string    `json:"name"`
	NbAlias                int       `json:"nb_alias"`
	RandomPrefixGeneration bool      `json:"random_prefix_generation"`
}

// DomainResponse represents the response for listing domains
type DomainResponse struct {
	CustomDomains []Domain `json:"custom_domains"`
}

type UpdateDomain struct {
	CatchAll               bool   `json:"catch_all" validate:"omitempty"`
	RandomPrefixGeneration bool   `json:"random_prefix_generation" validate:"omitempty"`
	Name                   string `json:"name" validate:"omitempty"`
	MailboxIds             []int  `json:"mailbox_ids" validate:"omitempty"`
}

type TrashAlias struct {
	Alias             string `json:"alias"`
	DeletionTimestamp int    `json:"deletion_timestamp"`
}

type TrashDomainResponse struct {
	Aliases []TrashAlias `json:"aliases"`
}

// GetDomains retrieves all domains for the user
func (c *Client) GetDomains() ([]Domain, error) {
	resp, err := c.doRequest(http.MethodGet, "/custom_domains", nil)
	if err != nil {
		return nil, err
	}

	var result DomainResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.CustomDomains, nil
}

// CreateDomain creates a new custom domain
func (c *Client) UpdateDomain(domainID int, updateDomain UpdateDomain) (*Domain, error) {
	endpoint := fmt.Sprintf("/custom_domains/%d", domainID)

	jsonData, err := json.Marshal(updateDomain)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal UpdateDomain data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPatch, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result Domain
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteDomain deletes a domain by ID
func (c *Client) GetDeletedAliasesDomain(domainID int) ([]TrashAlias, error) {
	endpoint := fmt.Sprintf("/custom_domains/%d/trash", domainID)

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result TrashDomainResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Aliases, nil
}
