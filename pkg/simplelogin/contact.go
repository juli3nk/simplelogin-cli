package simplelogin

import (
	"fmt"
	"net/http"
)

type ContactDeleteResponse struct {
	Deleted bool `json:"deleted"`
}

type ContactBlockResponse struct {
	BlockForward bool `json:"block_forward"`
}

// DeleteContact deletes a contact
func (c *Client) DeleteContact(contactID int) (*ContactDeleteResponse, error) {
	if contactID <= 0 {
		return nil, &ValidationError{Field: "contactID", Message: "contact ID must be positive"}
	}

	endpoint := fmt.Sprintf("/contacts/%d", contactID)

	resp, err := c.doRequest(http.MethodDelete, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result ContactDeleteResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ToggleContact blocks a contact
func (c *Client) ToggleContact(contactID int) (*ContactBlockResponse, error) {
	if contactID <= 0 {
		return nil, &ValidationError{Field: "contactID", Message: "contact ID must be positive"}
	}

	endpoint := fmt.Sprintf("/contacts/%d/toggle", contactID)

	resp, err := c.doRequest(http.MethodPatch, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result ContactBlockResponse
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
