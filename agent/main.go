package agent

import (
	sky "goskyline"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string

	httpClient *http.Client
}

func newClient(url *url.URL, token string) *Client {
	return &Client{
		BaseURL:    url,
		UserAgent:  "golang/sdkclient",
		Token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) ListHosts() ([]sky.Host, error) {
	return []sky.Host{}, nil
}
