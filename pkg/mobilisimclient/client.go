package mobilisimclient

import (
	"context"
	"errors"
	"vatansoft-sms-service/pkg/httpclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

type Client interface {
	HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error)

	OneToN(ctx context.Context, r model.RequestOneToN) (*model.ResourceOneToN, error)
}

type Config struct {
	RegisterURL string
	LoginURL    string
	SearchURL   string
	BackendURL  string
	APIKey      string
}

type client struct {
	registerURL string
	loginURL    string
	searchURL   string
	backendURL  string
	apiKey      string
	httpClient  httpclient.HTTPClient
}

func NewClient(c Config, hc httpclient.HTTPClient) Client {
	return &client{
		registerURL: c.RegisterURL,
		loginURL:    c.LoginURL,
		searchURL:   c.SearchURL,
		backendURL:  c.BackendURL,
		apiKey:      c.APIKey,
		httpClient:  hc,
	}
}

func (c *client) HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error) {
	resp, err := c.httpClient.HandleRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if c.httpClient.IsSuccessStatusCode(resp) {
		return resp, nil
	}

	//return nil, c.HandleException(resp)
	return nil, errors.New("test")
}
