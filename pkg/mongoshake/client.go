package mongoshake

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	alias   string
	client  *http.Client
}

type Option func(*Client)

func WithAlias(alias string) Option {
	return func(c *Client) {
		c.alias = alias
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.client.Timeout = timeout
	}
}

func NewClient(baseURL string, opts ...Option) *Client {
	c := &http.Client{
		// default timeout, 1s
		// override by WithTimeout
		Timeout: time.Second,
	}
	client := &Client{baseURL: baseURL, client: c}

	for _, opt := range opts {
		opt(client)
	}

	if client.alias == "" {
		client.alias = baseURL
	}

	return client
}

func (c *Client) GetAlias() string {
	return c.alias
}

func (c *Client) GetBaseURL() string {
	return c.baseURL
}

func (c *Client) GetRepl() (*Repl, error) {
	resp, err := c.client.Get(c.baseURL + "/repl")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repl Repl
	if err := json.NewDecoder(resp.Body).Decode(&repl); err != nil {
		return nil, err
	}

	return &repl, nil
}
