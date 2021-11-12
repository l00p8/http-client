package http_client

import (
	"time"

	"github.com/gojektech/heimdall"
	"github.com/gojektech/heimdall/hystrix"
)

// TODO Create client from a factory with different type of clients
func newHystrixClient(service string, cfg Config, client heimdall.Doer) *hystrix.Client {
	timeout := time.Duration(cfg.HttpTimeout) * time.Millisecond
	if client != nil {
		timeout = 0 * time.Millisecond
	}

	initTimeout := time.Duration(cfg.InitialTimeout) * time.Millisecond
	maxTimeout := time.Duration(cfg.MaxTimeout) * time.Millisecond
	maxJitter := time.Duration(cfg.MaxJitter) * time.Millisecond

	hystrixTimeout := time.Duration(cfg.HystrixTimeout) * time.Millisecond
	return hystrix.NewClient(
		hystrix.WithHTTPTimeout(timeout),
		hystrix.WithCommandName(service),
		hystrix.WithHystrixTimeout(hystrixTimeout),
		hystrix.WithMaxConcurrentRequests(cfg.MaxConcurrent),
		hystrix.WithErrorPercentThreshold(cfg.ErrorThreshold),
		hystrix.WithSleepWindow(cfg.Sleep),
		hystrix.WithRequestVolumeThreshold(cfg.RequestThreshold),
		hystrix.WithHTTPClient(client),
		hystrix.WithRetrier(heimdall.NewRetrier(heimdall.NewExponentialBackoff(
			initTimeout,
			maxTimeout,
			cfg.ExponentFactor,
			maxJitter,
		))),
		hystrix.WithRetryCount(cfg.RetryCount),
	)
}
