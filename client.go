package fluenceapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HostURL - Default Fluence API URL
const HostURL string = "https://api.fluence.dev"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	ApiKey     string
}

// NewClient -
func NewClient(host *string, apikey *string) (*Client, error) {
	// If API key not provided, return error
	if apikey == nil || *apikey == "" {
		return nil, errors.New("API key is required")
	}

	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
		ApiKey:     *apikey,
	}

	if host != nil {
		c.HostURL = *host
	}

	return &c, nil
}

// doRequest uses the client's API key
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.ApiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
