package agent

import (
	"net/http"
	"net/url"

	sky "github.com/linuxskyline/goskyline"
)goskyline

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string

	httpClient *http.Client
}

func NewClient(url *url.URL, token string) *Client {
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

func (c *Client) CreateUpdate() error {
	return nil
}