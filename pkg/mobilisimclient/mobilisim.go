package mobilisimclient

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"vatansoft-sms-service/pkg/httpclient"
	"vatansoft-sms-service/pkg/mobilisimclient/model"
)

const (
	toMessageURL = "text/advanced"
)

func (c *client) OneToN(ctx context.Context, r model.RequestOneToN) error {
	_, err := c.HandleRequest(ctx, httpclient.Request{
		URL:     c.MobilisimURL(toMessageURL),
		Method:  fiber.MethodGet,
		Body:    r,
		Headers: c.PrepareHeaders(),
	})
	//if err != nil {
	//	return nil, err
	//}

	//var res model.ResourceOneToN
	//
	//err = json.Unmarshal(resp.Body, &res)
	//if err != nil {
	//	return nil, err
	//}

	return err
}
