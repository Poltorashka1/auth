package config

import (
	"fmt"
)

// GRPCConfig interface
type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
	cert string
	key  string
}

// Address returns address of gRPC server
func (cfg *grpcConfig) Address() string {
	return fmt.Sprintf("%s:%s", cfg.host, cfg.port)
}

// NewGRPCConfig returns new GRPCConfig
func NewGRPCConfig() (GRPCConfig, error) {
	host, err := GetEnv("GRPC_HOST")
	if err != nil {
		return nil, err
	}
	port, err := GetEnv("GRPC_PORT")
	if err != nil {
		return nil, err
	}
	cert, err := GetEnv("GRPC_CERT")
	if err != nil {
		return nil, err
	}
	key, err := GetEnv("GRPC_KEY")
	if err != nil {
		return nil, err
	}
	return &grpcConfig{
		host: host,
		port: port,
		cert: cert,
		key:  key,
	}, nil
}
