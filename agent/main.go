package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	sky "github.com/linuxskyline/goskyline"
)

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

func (c *Client) CreateUpdate(update sky.Update) error {
	b, err := json.Marshal(update)
	if err != nil {
		return err
	}

	client := &http.Client{}
	r, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.BaseURL.String(), "updates"),
		bytes.NewBuffer(b),
	)
	r.Header.Set("HostToken", c.Token)
	r.Header.Set("Content-Type", "application/json")

	_, err = client.Do(r)
	if err != nil {
		return err
	}

	return nil
}
