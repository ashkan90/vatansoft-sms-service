package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

const (
	clientName     = "BiTaksi/1.0"
	defaultTimeout = 5 * time.Second
)

type Request struct {
	URL     string
	Method  string
	Body    interface{}
	Headers map[string]string

	Timeout time.Duration
}

type Response struct {
	StatusCode int
	Body       []byte
}

type ResponseErrorBag struct {
	Err   Error `json:"error"`
	Cause error `json:"cause"`

	Response
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (re ResponseErrorBag) Error() string {
	return fmt.Sprintf("status: %d - body: %s", re.StatusCode, string(re.Response.Body))
}

type HTTPClient interface {
	HandleRequest(ctx context.Context, req Request) (*Response, error)
	HandleException(resp *Response) error
	IsSuccessStatusCode(resp *Response) bool
	GetJSONHeaders() map[string]string
	IsCustomErrorType(errType string, err error) bool
}

type httpClient struct {
	client *fasthttp.Client
}

func NewHTTPClient() HTTPClient {
	fc := &fasthttp.Client{
		Name:                     clientName,
		NoDefaultUserAgentHeader: true,
	}

	return &httpClient{client: fc}
}

func (h *httpClient) HandleRequest(ctx context.Context, req Request) (*Response, error) {
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	//var externalSegment *newrelic.ExternalSegment
	//if tx := newrelic.FromContext(ctx); tx != nil {
	//	externalSegment = &newrelic.ExternalSegment{URL: req.URL}
	//	externalSegment.StartTime = tx.StartSegmentNow()
	//	defer externalSegment.End()
	//}

	request.SetRequestURI(req.URL)
	if req.Body != nil {
		body, bodyErr := h.prepareBody(req.Body)
		if bodyErr != nil {
			return nil, bodyErr
		}

		if body != nil {
			request.SetBody(body)
		}
	}

	request.Header.SetMethod(req.Method)
	for key, header := range req.Headers {
		request.Header.Set(key, header)
	}

	if req.Timeout <= 0 {
		req.Timeout = defaultTimeout
	}

	if err := h.client.DoTimeout(request, resp, req.Timeout); err != nil {
		return nil, fmt.Errorf("request err: %v", err)
	}

	respBody := resp.Body()
	respStatusCode := resp.StatusCode()

	var bytes []byte
	bytes = append(bytes, respBody...)

	//if externalSegment != nil {
	//	externalSegment.SetStatusCode(respStatusCode)
	//}

	return &Response{
		Body:       bytes,
		StatusCode: respStatusCode,
	}, nil
}

func (h *httpClient) HandleException(resp *Response) error {
	respErr := ResponseErrorBag{
		Response: Response{
			StatusCode: resp.StatusCode,
			Body:       resp.Body,
		},
	}

	if jsonErr := json.Unmarshal(resp.Body, &respErr); jsonErr != nil {
		respErr.Cause = fmt.Errorf("json err: %v", jsonErr)
	} else {
		respErr.Cause = errors.New("cause: nil")
	}

	return respErr
}

func (h *httpClient) IsSuccessStatusCode(resp *Response) bool {
	return resp.StatusCode >= fasthttp.StatusOK && resp.StatusCode < fasthttp.StatusMultipleChoices
}

func (h *httpClient) GetJSONHeaders() map[string]string {
	headers := make(map[string]string)
	//headers[utils.AcceptHeaderKey] = "application/json"
	//headers[utils.ContentTypeHeaderKey] = "application/json"

	return headers
}

func (h *httpClient) IsCustomErrorType(errType string, err error) bool {
	if errorBag, ok := err.(ResponseErrorBag); ok && errorBag.Err.Code == errType {
		return true
	}
	return false
}

func (h *httpClient) prepareBody(b interface{}) ([]byte, error) {
	if byteBody, byteBodyOk := b.([]byte); byteBodyOk {
		return byteBody, nil
	}

	body, encodeErr := json.Marshal(b)
	if encodeErr != nil {
		return nil, encodeErr
	}

	return body, nil
}
