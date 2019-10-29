package agent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/jsonapi"
	sky "github.com/linuxskyline/goskyline"
)

type Response struct {
}

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

	r, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s", c.BaseURL.String(), "updates"),
		bytes.NewBuffer(b),
	)
	r.Header.Set("HostToken", c.Token)
	r.Header.Set("Content-Type", "application/json")

	_, err = c.httpClient.Do(r)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUpdates() ([]*sky.Update, error) {
	r, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s", c.BaseURL.String(), "updates"),
		nil,
	)
	r.Header.Set("HostToken", c.Token)
	r.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}

	responseUpdates := []*sky.Update{}
	err = jsonapi.UnmarshalPayload(resp.Body, responseUpdates)
	if err != nil {
		return nil, err
	}

	return responseUpdates, nil
}

func (c *Client) DeleteUpdate(update *sky.Update) error {
	r, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/%s/%d", c.BaseURL.String(), "updates", update.ID),
		nil,
	)
	r.Header.Set("HostToken", c.Token)
	r.Header.Set("Content-Type", "application/json")

	_, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}

	return nil
}
