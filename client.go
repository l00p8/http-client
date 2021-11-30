package xclient

import (
	"context"
	"net/http"
	"time"

	"io"

	"github.com/gojek/heimdall/v7/hystrix"
)

type Client interface {
	Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (*http.Response, error)
}

func New(service string, cfg Config) Client {
	base := cfg.Transport
	if base == nil {
		base = http.DefaultTransport
	}
	httpClt := &http.Client{
		Transport: base,
		Timeout:   time.Duration(cfg.HttpTimeout) * time.Millisecond,
	}
	clt := newHystrixClient(service, cfg, httpClt)
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
