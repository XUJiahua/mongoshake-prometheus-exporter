package mongoshake

import (
	"encoding/json"
	"net/http"
	"time"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string) *Client {
	c := &http.Client{
		// refactor me: as an option of constructor
		Timeout: time.Second,
	}
	return &Client{baseURL: baseURL, client: c}
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
