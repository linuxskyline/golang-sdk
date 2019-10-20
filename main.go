package goskyline

import (
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string
	Token     string

	httpClient *http.Client
}

func (c *Client) ListHosts() ([]Host, error) {
	return []Host{}, nil
}
