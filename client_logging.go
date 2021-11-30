package xclient

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/l00p8/log"
)

func WithLogging(clt Client, logger log.Logger) Client {
	return &loggingClient{
		clt: clt,
		lg:  logger,
	}
}

type loggingClient struct {
	clt Client
	lg  log.Logger
}

func (clt *loggingClient) Request(ctx context.Context, method string, url string, body io.Reader, headers http.Header) (*http.Response, error) {
	now := time.Now()
	resp, err := clt.clt.Request(ctx, method, url, body, headers)
	clt.lg.Debug(method + " " + url + " " + time.Since(now).String())
	return resp, err
}
