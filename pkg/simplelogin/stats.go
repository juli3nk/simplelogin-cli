package simplelogin

import "net/http"

type Stats struct {
	NBAlias   int `json:"nb_alias"`
	NBBlock   int `json:"nb_block"`
	NBForward int `json:"nb_forward"`
	NBReply   int `json:"nb_reply"`
}

func (c *Client) GetStats() (*Stats, error) {
	endpoint := "/stats"

	resp, err := c.doRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	var result Stats
	if err := c.handleResponse(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
