package api

import (
	_ "auth/docs"
	"auth/internal/api/http/middleware"
	"auth/internal/api/http/router"
	apiUser "auth/internal/api/http/user"
	"auth/internal/config"
	"github.com/swaggo/http-swagger/v2"
	"net/http"
)

type Server struct {
	cfg         config.HTTPConfig
	router      router.Router
	userHandler apiUser.UserHandler
}

func (s *Server) Addr() string {
	return s.cfg.Address()
}

func NewServerHTTP(cfg config.HTTPConfig, router router.Router, userHandler apiUser.UserHandler) *Server {
	return &Server{
		cfg:         cfg,
		router:      router,
		userHandler: userHandler,
	}
}

func (s *Server) Start() error {
	return http.ListenAndServeTLS(s.Addr(), s.cfg.Cert(), s.cfg.Key(), s.router)
	// return http.ListenAndServe(s.Addr(), s.router)
}

func (s *Server) InitMiddleware() {
	s.router.AddMiddleware(middleware.Cors)
}

func (s *Server) InitRoutes() {
	s.router.AddRoute(http.MethodGet, "/swagger/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:50052/swagger/doc.json")))
	s.router.AddRoute(http.MethodPost, "/signup", s.userHandler.SignUp)
	s.router.AddRoute(http.MethodGet, "/users", middleware.CheckAccess(s.userHandler.Users))
	s.router.AddRoute(http.MethodGet, "/verify", s.userHandler.EmailVerify) // todo mb post
	s.router.AddRoute(http.MethodPost, "/signin", s.userHandler.SignIn)
	// todo check how it work and after that delete mid
	s.router.AddRoute(http.MethodGet, "/refresh", middleware.CheckAccess(s.userHandler.GetAccessToken))
	s.router.AddRoute(http.MethodDelete, "/signout", s.userHandler.SignOut)
	s.router.AddRoute(http.MethodGet, "/check-access", s.userHandler.CheckAccess)
	// todo ChangePassword
	// todo UpdateUser
	// todo DeleteUser
	// todo GetUsers
}
