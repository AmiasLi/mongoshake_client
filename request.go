package mongoshake

import "github.com/go-resty/resty/v2"

func (c *Client) request(url string, p interface{}) (interface{}, error) {
	restClient := resty.New().SetBaseURL(c.BaseURL)

	_, err := restClient.R().
		SetResult(p).
		ForceContentType("application/json").
		Get(url)

	if err != nil {
		return nil, err
	}

	return p, nil
}
