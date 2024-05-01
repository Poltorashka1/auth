package app

import (
	apiUserGRPC "auth/internal/api/gRPC/user"
	"auth/internal/api/http/router"
	apiUserHTTP "auth/internal/api/http/user"
	"auth/internal/closer"
	"auth/internal/config"
	"auth/internal/db"
	"auth/internal/db/pg"
	cache "auth/internal/db/redis"
	"auth/internal/logger"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/internal/smtp"
	"context"
)

type provider struct {
	log logger.Logger

	pgConfig    config.PostgresConfig
	grpcConfig  config.GRPCConfig
	httpConfig  config.HTTPConfig
	smtpConfig  config.SMTPConfig
	redisConfig config.RedisConfig

	userAPIgRPC *apiUserGRPC.Implementation // todo rename
	userAPIHTTP *apiUserHTTP.UserHandler
	service     service.Service
	repository  repository.Repository

	// pgPool         db.Client
	pgPool     db.DB
	cache      db.Cache
	smtpClient smtp.SMTP

	router router.Router
}

func newProvider() *provider {
	return &provider{}
}

func (p *provider) Logger() logger.Logger {
	if p.log == nil {
		l := logger.Load()
		p.log = l
	}
	return p.log
}

func (p *provider) RedisConfig() config.RedisConfig {
	if p.redisConfig == nil {
		cfg := config.NewRedisConfig()

		p.redisConfig = cfg
	}
	return p.redisConfig
}

func (p *provider) HTTPConfig() config.HTTPConfig {
	const op = "app.HttpConfig"

	if p.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			p.Logger().FatalOp(op, err)
		}

		p.httpConfig = cfg
	}

	return p.httpConfig
}

func (p *provider) SMTPConfig() config.SMTPConfig {
	const op = "app.SMTPConfig"
	if p.smtpConfig == nil {
		cfg, err := config.NewSMTPConfig()
		if err != nil {
			p.Logger().FatalOp(op, err)
		}
		p.smtpConfig = cfg
	}
	return p.smtpConfig
}

func (p *provider) PGConfig() config.PostgresConfig {
	const op = "app.PGConfig"

	if p.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			p.Logger().FatalOp(op, err)
		}

		p.pgConfig = cfg
	}
	return p.pgConfig

}

func (p *provider) GRPCConfig() config.GRPCConfig {
	const op = "app.GRPCConfig"

	if p.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			p.Logger().FatalOp(op, err)
		}

		p.grpcConfig = cfg
	}
	return p.grpcConfig
}
func (p *provider) SMTPClient() smtp.SMTP {
	const op = "app.SMTPClient"

	if p.smtpClient == nil {
		smtpClient, err := smtp.New(p.SMTPConfig(), p.Logger())
		if err != nil {
			p.Logger().FatalOp(op, err)
		}

		p.smtpClient = smtpClient

		closer.Add(smtpClient.Close)
	}
	return p.smtpClient
}

func (p *provider) Cache() db.Cache {
	if p.cache == nil {
		c := cache.New(p.RedisConfig())
		p.cache = c

		closer.Add(c.Close)
	}
	return p.cache
}

func (p *provider) PgPool(ctx context.Context) db.DB { // db.Client
	const op = "app.PgPool"

	if p.pgPool == nil {
		client, err := pg.NewDB(ctx, p.PGConfig().DSN(), p.Logger())
		if err != nil {
			p.Logger().FatalOp(op, err)
		}

		if err := client.Ping(ctx); err != nil {
			p.Logger().FatalOp(op, err)
		}

		p.pgPool = client
		closer.Add(client.Close)
	}

	return p.pgPool
}

func (p *provider) Repository(ctx context.Context) repository.Repository {
	if p.repository == nil {
		repo := repository.New(p.PgPool(ctx), p.Cache(), p.Logger())
		p.repository = repo
	}
	return p.repository
}

func (p *provider) Service(ctx context.Context) service.Service {
	if p.service == nil {
		p.service = service.New(p.Repository(ctx), p.PgPool(ctx), p.SMTPClient(), p.Logger())
	}
	return p.service
}

func (p *provider) UserAPIHTTP(ctx context.Context) *apiUserHTTP.UserHandler {
	if p.userAPIHTTP == nil {
		h := apiUserHTTP.New(p.Service(ctx), p.Logger(), p.HTTPConfig().JwtSecret())
		p.userAPIHTTP = h
	}

	return p.userAPIHTTP
}

func (p *provider) UserAPIGRPC(ctx context.Context) *apiUserGRPC.Implementation {
	if p.userAPIgRPC == nil {
		a := apiUserGRPC.New(p.Service(ctx), p.Logger())

		p.userAPIgRPC = &a
	}

	return p.userAPIgRPC
}

func (p *provider) Router() router.Router {
	if p.router == nil {
		rout := router.New()
		p.router = rout
	}

	return p.router
}
