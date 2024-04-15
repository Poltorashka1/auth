package app

import (
	httpserver "auth/internal/api/http"
	"auth/internal/closer"
	"auth/internal/config"
	"auth/pkg"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
	"sync"
)

type App struct {
	wg       sync.WaitGroup
	provider *provider
	grpcServ *grpc.Server
	httpServ *httpserver.Server
}

func New(ctx context.Context) *App {
	a := &App{}
	a.initDeps(ctx)
	return a
}

// Run starts grpc server and http server
func (a *App) Run() {
	defer func() {
		closer.CloseAll()
	}()

	HTTPEnd := make(chan error)
	gRPCEnd := make(chan error)

	a.wg.Add(2)
	go func() {
		defer a.wg.Done()
		err := a.startHTTPServer()
		if err != nil {
			HTTPEnd <- err
		}
	}()
	go func() {
		defer a.wg.Done()
		err := a.runGRPCServer()
		if err != nil {
			gRPCEnd <- err
		}
	}()

	go func() {
		a.wg.Wait()
		close(HTTPEnd)
		close(gRPCEnd)
	}()

	// todo из за использования closer.CloseAll() не можем вернуть ошибку так как он прекращает работу программы.
	for {
		select {
		case err := <-gRPCEnd:
			a.provider.Logger().Error(err)
		case err := <-HTTPEnd:
			a.provider.Logger().Error(err)
		}
	}
}

func (a *App) initDeps(ctx context.Context) {
	a.loadConfig()

	a.initProvider()

	a.initGrpc(ctx)
	a.initHTTP(ctx)
}

func (a *App) loadConfig() {
	config.Load(".env")
}

func (a *App) initProvider() {
	a.provider = newProvider()
}

func (a *App) initHTTP(ctx context.Context) {
	// create server with router and handler
	a.httpServ = httpserver.NewServerHTTP(a.provider.HTTPConfig(), a.provider.Router(), *a.provider.UserAPIHTTP(ctx))
	// add middleware to router
	a.httpServ.InitMiddleware()
	// connect handler to router
	a.httpServ.InitRoutes()
}

func (a *App) startHTTPServer() error {
	a.provider.Logger().Info("start https server on " + a.provider.HTTPConfig().Address())
	err := a.httpServ.Start()
	if err != nil {
		return err
	}
	return nil
}
func (a *App) initGrpc(ctx context.Context) {
	const op = "app.initGrpc"

	creds, err := credentials.NewServerTLSFromFile(a.provider.HTTPConfig().Cert(), a.provider.HTTPConfig().Key())
	if err != nil {
		a.provider.Logger().FatalOp(op, err)
	}
	a.grpcServ = grpc.NewServer(grpc.Creds(creds))
	reflection.Register(a.grpcServ)
	pkg.RegisterAuthServer(a.grpcServ, a.provider.UserAPIGRPC(ctx))
}

func (a *App) runGRPCServer() error {
	l, err := net.Listen("tcp", a.provider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	a.provider.Logger().Info("start grpc server on " + a.provider.grpcConfig.Address())

	err = a.grpcServ.Serve(l)
	if err != nil {
		return err
	}

	return nil
}
