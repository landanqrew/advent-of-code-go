package client

import (
	"bytes"
	"net/http"
)


type HttpClient struct {
	Client *http.Client
	BaseURL string
}

func NewHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		Client: &http.Client{},
		BaseURL: baseURL,
	}
}

func (c *HttpClient) Get(path string) (*http.Response, error) {
	url := c.BaseURL + path
	return c.Client.Get(url)
}

func (c *HttpClient) Post(path string, body []byte) (*http.Response, error) {
	url := c.BaseURL + path
	return c.Client.Post(url, "application/json", bytes.NewBuffer(body))
}