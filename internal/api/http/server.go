package api

import (
	_ "auth/docs"
	"auth/internal/api/http/middleware"
	"auth/internal/api/http/router"
	apiUser "auth/internal/api/http/user"
	"auth/internal/config"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
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
	s.router.AddRoute(http.MethodGet, "/user", s.Authorization(s.userHandler.GetUser))
	s.router.AddRoute(http.MethodGet, "/verify", s.userHandler.EmailVerify) // todo mb post
	s.router.AddRoute(http.MethodPost, "/signin", s.userHandler.SignIn)
	s.router.AddRoute(http.MethodGet, "/refresh", s.userHandler.RefreshToken)
}

func (s *Server) Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("invalid token"))
			return
		}

		// todo uznat kak eto work
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(s.cfg.JwtSecret()), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("invalid token"))
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("invalid token"))
			return
		}

		next.ServeHTTP(w, r)
	}
}
