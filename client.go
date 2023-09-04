package mongoshake

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	BaseURL string
}

func NewClient(url string) *Client {
	return &Client{BaseURL: url}
}

func (c *Client) VerifyURL(url string) error {
	restClient := resty.New().SetBaseURL(url)
	if _, err := restClient.R().Get("/conf"); err != nil {
		return fmt.Errorf("url is invalid")
	}
	return nil
}

func New(url string) (*Client, error) {
	if url == "" {
		return nil, fmt.Errorf("url is empty")
	}

	client := NewClient(url)
	if err := client.VerifyURL(url); err != nil {
		return nil, err
	}

	return client, nil
}
