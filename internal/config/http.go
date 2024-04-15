package config

import "fmt"

type HTTPConfig interface {
	Address() string
	Port() string
	Host() string
	Cert() string
	Key() string
	JwtSecret() string
}

type httpConfig struct {
	host      string
	port      string
	cert      string
	key       string
	jwtSecret string
}

func NewHTTPConfig() (HTTPConfig, error) {
	host, err := GetEnv("HTTP_HOST")
	if err != nil {
		return nil, err
	}
	port, err := GetEnv("HTTP_PORT")
	if err != nil {
		return nil, err
	}
	cert, err := GetEnv("HTTP_CERT")
	if err != nil {
		return nil, err
	}
	key, err := GetEnv("HTTP_KEY")
	if err != nil {
		return nil, err
	}
	jwtSecret, err := GetEnv("JWT_SECRET")
	if err != nil {
		return nil, err
	}
	return &httpConfig{
		host:      host,
		port:      port,
		cert:      cert,
		key:       key,
		jwtSecret: jwtSecret,
	}, nil
}

func (cfg *httpConfig) Cert() string {
	return cfg.cert
}

func (cfg *httpConfig) Key() string {
	return cfg.key
}

func (cfg *httpConfig) Address() string {
	return fmt.Sprintf("%s:%s", cfg.host, cfg.port)
}

func (cfg *httpConfig) Port() string {
	return cfg.port
}

func (cfg *httpConfig) Host() string {
	return cfg.host
}

func (cfg *httpConfig) JwtSecret() string {
	return cfg.jwtSecret
}
