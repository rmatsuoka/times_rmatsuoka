package xhttp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Client struct {
	Client HTTPClient
}

var DefaultClient = &Client{}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	client := c.Client
	if client == nil {
		client = http.DefaultClient
	}
	return client.Do(req)
}

func (c *Client) GetJSON(ctx context.Context, url string, resbody any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return err
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		io.Copy(io.Discard, res.Body)
		res.Body.Close()
	}()

	buf, err := io.ReadAll(io.LimitReader(res.Body, 1<<20))
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, resbody)
}
