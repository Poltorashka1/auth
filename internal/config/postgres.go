package config

import "fmt"

type PostgresConfig interface {
	DSN() string
}

type postgresConfig struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func NewPGConfig() (PostgresConfig, error) {
	host, err := GetEnv("PG_HOST")
	if err != nil {
		return nil, err
	}
	port, err := GetEnv("PG_PORT")
	if err != nil {
		return nil, err
	}
	user, err := GetEnv("PG_USER")
	if err != nil {
		return nil, err
	}
	password, err := GetEnv("PG_PASSWORD")
	if err != nil {
		return nil, err
	}
	dbname, err := GetEnv("PG_DATABASE_NAME")
	if err != nil {
		return nil, err
	}
	return &postgresConfig{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		dbname:   dbname,
	}, nil
}

func (cfg *postgresConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.user, cfg.password, cfg.host, cfg.port, cfg.dbname)
}
