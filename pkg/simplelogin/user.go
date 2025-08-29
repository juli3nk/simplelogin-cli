package simplelogin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserInfo struct {
	Name              string `json:"name"`
	IsPremium         bool   `json:"is_premium"`
	Email             string `json:"email"`
	InTrial           bool   `json:"in_trial"`
	ProfilePictureURL string `json:"profile_picture_url"`
	MaxAliasFreePlan  int    `json:"max_alias_free_plan"`
}

type UserInfoUpdate struct {
	ProfilePicture string `json:"profile_picture"`
	Name           string `json:"name"`
}

func (c *Client) GetUserInfo() (*UserInfo, error) {
	endpoint := "/user_info"

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result UserInfo
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) UpdateUserInfo(userInfo UserInfoUpdate) (*UserInfo, error) {
	endpoint := "/user_info"

	jsonData, err := json.Marshal(userInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal UserInfoUpdate data: %w", err)
	}

	resp, err := c.doRequest(http.MethodPut, endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	var result UserInfo
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
