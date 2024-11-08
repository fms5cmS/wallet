package client

import "time"

type EndpointType uint8

const (
	EndpointTypeUnknown EndpointType = iota
	EndpointTypeRPC
	EndpointTypeWS
	EndpointTypeHTTP
)

type Config struct {
	Endpoints           map[EndpointType]string
	HTTTPHeaders        map[string]string
	NativeTokenSymbol   string
	NativeTokenDecimals uint8
	Timeout             time.Duration
}

type ConfigOption func(*Config)

func WithConfigTimeoutOption(timeout time.Duration) ConfigOption {
	return func(cfg *Config) {
		cfg.Timeout = timeout
	}
}

func WithConfigEndpointsOption(endpoints map[EndpointType]string) ConfigOption {
	return func(cfg *Config) {
		cfg.Endpoints = endpoints
	}
}

func WithConfigHTTPHeadersOption(headers map[string]string) ConfigOption {
	return func(cfg *Config) {
		cfg.HTTTPHeaders = headers
	}
}

func WithConfigNativeTokenOption(symbol string, decimals uint8) ConfigOption {
	return func(cfg *Config) {
		cfg.NativeTokenSymbol = symbol
		cfg.NativeTokenDecimals = decimals
	}
}
