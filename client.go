package http_client

import (
	"context"
	"net/http"

	"io"

	"github.com/gojektech/heimdall/hystrix"
)

type Client interface {
	Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (*http.Response, error)
}

func New(service string, cfg Config) Client {
	clt := newHystrixClient(service, cfg, nil)
	return &defaultClient{
		clt: clt,
	}
}

type defaultClient struct {
	clt *hystrix.Client
}

func (clt *defaultClient) Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if headers == nil {
		headers = make(http.Header)
	}
	req.Header = headers

	resp, err := clt.clt.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
