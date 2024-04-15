package config

import "fmt"

type SMTPConfig interface {
	Address() string
	Host() string
	From() string
	Username() string
	Password() string
	Template() string
}
type smtpConfig struct {
	host     string
	port     string
	username string
	password string
	from     string
	template string
}

func NewSMTPConfig() (SMTPConfig, error) {
	host, err := GetEnv("SMTP_HOST")
	if err != nil {
		return nil, err
	}
	port, err := GetEnv("SMTP_PORT")
	if err != nil {
		return nil, err
	}
	username, err := GetEnv("SMTP_USER")
	if err != nil {
		return nil, err
	}
	password, err := GetEnv("SMTP_PASS")
	if err != nil {
		return nil, err
	}
	from, err := GetEnv("SMTP_USER")
	if err != nil {
		return nil, err
	}
	template, err := GetEnv("SMTP_TEMPLATE")
	if err != nil {
		return nil, err
	}
	return &smtpConfig{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
		template: template,
	}, nil
}

func (cfg *smtpConfig) Address() string {
	return fmt.Sprintf("%s:%s", cfg.host, cfg.port)
}

func (cfg *smtpConfig) Host() string {
	return cfg.host
}

func (cfg *smtpConfig) From() string {
	return cfg.from
}

func (cfg *smtpConfig) Template() string {
	return cfg.template
}

func (cfg *smtpConfig) Username() string {
	return cfg.username
}

func (cfg *smtpConfig) Password() string {
	return cfg.password
}
