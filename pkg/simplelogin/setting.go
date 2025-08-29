package simplelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Setting struct {
	AliasGenerator           string `json:"alias_generator" validate:"omitempty,oneof=word uuid"`
	Notification             bool   `json:"notification"`
	RandomAliasDefaultDomain string `json:"random_alias_default_domain" validate:"omitempty"`
	SenderFormat             string `json:"sender_format" validate:"omitempty,oneof=AT A NAME_ONLY AT_ONLY NO_NAME"`
	RandomAliasSuffix        string `json:"random_alias_suffix" validate:"omitempty,oneof=word random_string"`
}

type SettingDomain struct {
	Domain   string `json:"domain"`
	IsCustom bool   `json:"is_custom"`
}

// ValidateSetting validates a Setting struct (all fields required)
func (s *Setting) Validate(availableDomains []SettingDomain) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Additional validation for domain
	if !isValidDomain(s.RandomAliasDefaultDomain, availableDomains) {
		domainNames := getDomainNames(availableDomains)
		return fmt.Errorf("random_alias_default_domain must be one of: %s", strings.Join(domainNames, ", "))
	}

	return nil
}

// ValidatePartial validates a Setting struct for partial updates (all fields optional)
func (s *Setting) ValidatePartial(availableDomains []SettingDomain) error {
	validate := validator.New()

	if err := validate.Struct(s); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}

	// Additional validation for domain only if it's provided
	if s.RandomAliasDefaultDomain != "" && !isValidDomain(s.RandomAliasDefaultDomain, availableDomains) {
		domainNames := getDomainNames(availableDomains)
		return fmt.Errorf("random_alias_default_domain must be one of: %s", strings.Join(domainNames, ", "))
	}

	return nil
}

// isValidDomain checks if a domain exists in the available domains
func isValidDomain(domain string, availableDomains []SettingDomain) bool {
	for _, d := range availableDomains {
		if d.Domain == domain {
			return true
		}
	}
	return false
}

// getDomainNames returns a slice of domain names
func getDomainNames(domains []SettingDomain) []string {
	names := make([]string, len(domains))
	for i, domain := range domains {
		names[i] = domain.Domain
	}
	return names
}

func (c *Client) GetSetting() (*Setting, error) {
	endpoint := "/setting"

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result Setting
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateSetting(setting Setting) (*Setting, error) {
	// Get available domains for validation
	availableDomains, err := c.GetSettingDomains()
	if err != nil {
		return nil, fmt.Errorf("failed to get available domains for validation: %w", err)
	}

	// Validate the setting for partial update (all fields optional)
	if err := setting.ValidatePartial(availableDomains); err != nil {
		return nil, fmt.Errorf("setting validation failed: %w", err)
	}

	endpoint := "/setting"

	jsonData, err := json.Marshal(setting)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Setting data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result Setting
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetSettingDomains() ([]SettingDomain, error) {
	endpoint := "/v2/setting/domains"

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result []SettingDomain
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}
