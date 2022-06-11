package mobilisimclient

import (
	"context"
	"errors"
	"fmt"
	"vatansoft-sms-service/pkg/httpclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

type Client interface {
	HandleRequest(ctx context.Context, req httpclient.Request) (*httpclient.Response, error)

	PrepareHeaders() map[string]string
	MobilisimURL(uri string) string

	OneToN(ctx context.Context, r model.RequestOneToN) (*model.ResourceOneToN, error)
}

type Config struct {
	MobilisimURL string
	APIKey       string
}

type client struct {
	mobilisimURL string
	apiKey       string
	httpClient   httpclient.HTTPClient
}

func NewClient(c Config, hc httpclient.HTTPClient) Client {
	return &client{
		mobilisimURL: c.MobilisimURL,
		apiKey:       c.APIKey,
		httpClient:   hc,
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
	return nil, errors.New("code there")
}

func (c *client) MobilisimURL(uri string) string {
	return fmt.Sprintf("%s/sms/1/%s", c.mobilisimURL, uri)
}

func (c *client) PrepareHeaders() map[string]string {
	return map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Basic " + c.apiKey,
	}
}
