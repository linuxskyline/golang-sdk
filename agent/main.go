package agent

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

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

	responseUpdates, err := jsonapi.UnmarshalManyPayload(resp.Body, reflect.TypeOf(new(sky.Update)))
	if err != nil {
		return nil, err
	}

	updates := []*sky.Update{}
	for _, update := range responseUpdates {
		u, ok := update.(*sky.Update)
		if !ok {
			return nil, errors.New("error converting response type to update")
		}

		updates = append(updates, u)
	}

	return updates, nil
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
