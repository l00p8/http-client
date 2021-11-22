package xclient

import "net/http"

type Config struct {
	HttpTimeout int `envconfig:"http_timeout" default:"300"` // 300ms

	// Hystrix Circuit Breaker strategy
	HystrixTimeout   int `envconfig:"hystrix_timeout" default:"1100"` // 1.1s
	MaxConcurrent    int `envconfig:"max_concurrent" default:"100"`
	ErrorThreshold   int `envconfig:"error_threshold" default:"25"`
	Sleep            int `envconfig:"sleep" default:"10"` // sleep window
	RequestThreshold int `envconfig:"request_threshold" default:"10"`

	// Exponential Backoff Retry strategy
	InitialTimeout int     `envconfig:"initial_timeout" defualt:"2"` // 2ms
	MaxTimeout     int     `envconfig:"max_timeout" default:"5000"`  // 1s
	ExponentFactor float64 `envconfig:"exponent_factor" default:"2"` // Multiplier
	MaxJitter      int     `envconfig:"max_jitter" default:"2"`      // 2ms (it should be more than 1ms)
	RetryCount     int     `envconfig:"retry_count" default:"4"`

	Transport http.RoundTripper
}

func DefaultConfig() Config {
	return Config{
		HttpTimeout: 2500,
		// Hystrix Circuit Breaker strategy
		HystrixTimeout:   5100,
		MaxConcurrent:    100,
		ErrorThreshold:   25,
		Sleep:            10,
		RequestThreshold: 10,
		// Exponential Backoff Retry strategy
		InitialTimeout: 2,
		MaxTimeout:     10000,
		ExponentFactor: 2,
		MaxJitter:      2,
		RetryCount:     1,
		Transport:      http.DefaultTransport,
	}
}
