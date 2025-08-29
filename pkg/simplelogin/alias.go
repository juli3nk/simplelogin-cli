package simplelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// AliasOptions represents available options for creating aliases
type AliasOptions struct {
	CanCreate        bool                 `json:"can_create"`
	PrefixSuggestion string               `json:"prefix_suggestion"`
	Suffixes         []AliasOptionsSuffix `json:"suffixes"`
}

type AliasOptionsSuffix struct {
	SignedSuffix string `json:"signed_suffix"`
	Suffix       string `json:"suffix"`
	IsCustom     bool   `json:"is_custom"`
	IsPremium    bool   `json:"is_premium"`
}

type AliasCreateCustomOptions struct {
	Hostname     string `json:"hostname"`
	AliasPrefix  string `json:"alias_prefix"`
	SignedSuffix string `json:"signed_suffix"`
	MailboxIDs   []int  `json:"mailbox_ids"`
	Note         string `json:"note"`
	Name         string `json:"name"`
}

type AliasCreateRandomOptions struct {
	Hostname string `json:"hostname"`
	Mode     string `json:"mode"`
	Note     string `json:"note"`
}

// Alias represents a SimpleLogin alias
type Alias struct {
	CreationDate      string        `json:"creation_date"`
	CreationTimestamp int           `json:"creation_timestamp"`
	Email             string        `json:"email"`
	Name              string        `json:"name"`
	Enabled           bool          `json:"enabled"`
	ID                int           `json:"id"`
	Mailbox           Mailbox       `json:"mailbox"`
	Mailboxes         []Mailbox     `json:"mailboxes"`
	LatestActivity    AliasActivity `json:"latest_activity"`
	NbBlock           int           `json:"nb_block"`
	NbForward         int           `json:"nb_forward"`
	NbReply           int           `json:"nb_reply"`
	Note              string        `json:"note"`
	Pinned            bool          `json:"pinned"`
}

type AliasListOptions struct {
	Pinned   bool   `json:"pinned"`   // (Optional) If set, only pinned aliases are returned.
	Disabled bool   `json:"disabled"` // (Optional) If set, only disabled aliases are returned.
	Enabled  bool   `json:"enabled"`  // (Optional) If set, only enabled aliases are returned. Please note pinned, disabled, enabled are exclusive, i.e. only one can be present.
	Query    string `json:"query"`    // (Optional) Included in request body. Some frameworks might prevent GET request having a non-empty body, in this case this endpoint also supports POST.
}

// AliasResponse represents the response for listing aliases
type AliasResponse struct {
	Aliases []Alias `json:"aliases"`
}

type AliasDeleteResponse struct {
	Deleted bool `json:"deleted"`
}

// ToggleResponse represents the response for toggling an alias
type AliasToggleResponse struct {
	Enabled bool `json:"enabled"`
}

// Activity represents a SimpleLogin alias activity
type AliasActivity struct {
	Action              string `json:"action"`
	From                string `json:"from"`
	Timestamp           int    `json:"timestamp"`
	To                  string `json:"to"`
	ReverseAlias        string `json:"reverse_alias"`
	ReverseAliasAddress string `json:"reverse_alias_address"`
}

// ActivityResponse represents the response for listing activities
type AliasActivitiesResponse struct {
	Activities []AliasActivity `json:"activities"`
}

type AliasUpdateOptions struct {
	Note       string `json:"note"`
	MailboxID  int    `json:"mailbox_id"`
	Name       string `json:"name"`
	MailboxIDs []int  `json:"mailbox_ids"`
	DisablePGP bool   `json:"disable_pgp"`
	Pinned     bool   `json:"pinned"`
}

type AliasContact struct {
	ID                     int    `json:"id"`
	Contact                string `json:"contact"`
	CreationDate           string `json:"creation_date"`
	CreationTimestamp      int    `json:"creation_timestamp"`
	LastEmailSentDate      string `json:"last_email_sent_date"`
	LastEmailSentTimestamp int    `json:"last_email_sent_timestamp"`
	ReverseAlias           string `json:"reverse_alias"`
	BlockForward           bool   `json:"block_forward"`
}

type AliasContactResponse struct {
	Contacts []AliasContact `json:"contacts"`
}

type AliasContactCreateResponse struct {
	AliasContact

	Existed bool `json:"existed"`
}

// GetAliasOptions retrieves available options for creating aliases
func (c *Client) GetAliasOptions(hostname string) (*AliasOptions, error) {
	endpoint := "/v5/alias/options"

	if hostname != "" {
		endpoint = fmt.Sprintf("/v5/alias/options?hostname=%s", url.QueryEscape(hostname))
	}

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result AliasOptions
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateCustomAlias creates a new custom alias
func (c *Client) CreateCustomAlias(hostname string, options AliasCreateCustomOptions) (*Alias, error) {
	endpoint := "/v3/alias/custom/new"

	if hostname != "" {
		endpoint = fmt.Sprintf("%s?hostname=%s", endpoint, url.QueryEscape(hostname))
	}

	jsonData, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal create data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result Alias
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateRandomAlias creates a new random alias
func (c *Client) CreateRandomAlias(hostname, mode, note string) (*Alias, error) {
	endpoint := "/alias/random/new"

	urlValues := url.Values{}
	if hostname != "" {
		urlValues.Add("hostname", hostname)
	}
	if mode != "" {
		urlValues.Add("mode", mode)
	}

	if len(urlValues) > 0 {
		endpoint = fmt.Sprintf("%s?%s", endpoint, urlValues.Encode())
	}

	jsonData, err := json.Marshal(map[string]string{"note": note})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal create data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result Alias
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAliases retrieves all aliases for the user
func (c *Client) GetAliases(options AliasListOptions, pageID int) ([]Alias, error) {
	endpoint := fmt.Sprintf("/v2/aliases?page_id=%d", pageID)

	urlValues := url.Values{}
	if options.Pinned {
		urlValues.Add("pinned", "true")
	}
	if options.Disabled {
		urlValues.Add("disabled", "true")
	}
	if options.Enabled {
		urlValues.Add("enabled", "true")
	}
	if options.Query != "" {
		urlValues.Add("query", options.Query)
	}

	if len(urlValues) > 0 {
		endpoint = fmt.Sprintf("%s&%s", endpoint, urlValues.Encode())
	}

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result AliasResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Aliases, nil
}

// GetAllAliases retrieves all aliases across all pages
func (c *Client) GetAllAliases(options AliasListOptions) ([]Alias, error) {
	var allAliases []Alias
	pageID := 0

	for {
		aliases, err := c.GetAliases(options, pageID)
		if err != nil {
			return nil, err
		}

		if len(aliases) == 0 {
			break
		}

		allAliases = append(allAliases, aliases...)
		pageID++
	}

	return allAliases, nil
}

// GetAlias retrieves a specific alias by ID
func (c *Client) GetAlias(aliasID int) (*Alias, error) {
	endpoint := fmt.Sprintf("/aliases/%d", aliasID)

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result Alias
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteAlias deletes an alias by ID
func (c *Client) DeleteAlias(aliasID int) (bool, error) {
	endpoint := fmt.Sprintf("/aliases/%d", aliasID)

	resp, err := c.doRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return false, err
	}

	var result AliasDeleteResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return false, err
	}

	return result.Deleted, nil
}

// ToggleAlias enables or disables an alias
func (c *Client) ToggleAlias(aliasID int) (*AliasToggleResponse, error) {
	endpoint := fmt.Sprintf("/aliases/%d/toggle", aliasID)

	resp, err := c.doRequest(http.MethodPost, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result AliasToggleResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// GetAliasActivities retrieves all activities for a specific alias
func (c *Client) GetAliasActivities(aliasID, pageID int) ([]AliasActivity, error) {
	endpoint := fmt.Sprintf("/aliases/%d/activities?page_id=%d", aliasID, pageID)

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result []AliasActivity
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// GetAllAliasActivities retrieves all activities for a specific alias across all pages
func (c *Client) GetAllAliasActivities(aliasID int) ([]AliasActivity, error) {
	var allActivities []AliasActivity
	pageID := 0

	for {
		activities, err := c.GetAliasActivities(aliasID, pageID)
		if err != nil {
			return nil, err
		}

		if len(activities) == 0 {
			break
		}

		allActivities = append(allActivities, activities...)
		pageID++
	}

	return allActivities, nil
}

// UpdateAlias updates an alias's information
func (c *Client) UpdateAlias(aliasID int, options AliasUpdateOptions) error {
	endpoint := fmt.Sprintf("/aliases/%d", aliasID)

	jsonData, err := json.Marshal(options)
	if err != nil {
		return fmt.Errorf("failed to marshal update data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPatch, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	if err := c.handleResponse(resp, nil); err != nil {
		return err
	}

	return nil
}

// GetAliasContacts retrieves contacts for a specific alias with pagination
func (c *Client) GetAliasContacts(aliasID, pageID int) ([]AliasContact, error) {
	endpoint := fmt.Sprintf("/aliases/%d/contacts?page_id=%d", aliasID, pageID)

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result AliasContactResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return result.Contacts, nil
}

// GetAllAliasContacts retrieves all contacts for a specific alias across all pages
func (c *Client) GetAllAliasContacts(aliasID int) ([]AliasContact, error) {
	var allContacts []AliasContact
	pageID := 0

	for {
		contacts, err := c.GetAliasContacts(aliasID, pageID)
		if err != nil {
			return nil, err
		}

		if len(contacts) == 0 {
			break
		}

		allContacts = append(allContacts, contacts...)
		pageID++
	}

	return allContacts, nil
}

// CreateAliasContact creates a new contact for an alias
func (c *Client) CreateAliasContact(aliasID int, contact string) (*AliasContactCreateResponse, error) {
	endpoint := fmt.Sprintf("/aliases/%d/contacts", aliasID)

	jsonData, err := json.Marshal(map[string]string{"contact": contact})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal contact data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPost, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result AliasContactCreateResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
