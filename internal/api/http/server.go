package api

import (
	_ "auth/docs"
	apiJson "auth/internal/api/http/json"
	"auth/internal/api/http/middleware"
	"auth/internal/api/http/response"
	"auth/internal/api/http/router"
	apiUser "auth/internal/api/http/user"
	"auth/internal/config"
	"fmt"
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
	s.router.AddRoute(http.MethodGet, "/users", s.checkAccess(s.userHandler.GetUser))
	s.router.AddRoute(http.MethodGet, "/verify", s.userHandler.EmailVerify) // todo mb post
	s.router.AddRoute(http.MethodPost, "/signin", s.userHandler.SignIn)
	s.router.AddRoute(http.MethodGet, "/refresh", s.userHandler.GetAccessToken)
	s.router.AddRoute(http.MethodDelete, "/signout", s.userHandler.SignOut)
	s.router.AddRoute(http.MethodPost, "/check-access", s.userHandler.CheckAccess)
	// todo ChangePassword
	// todo UpdateUser
	// todo DeleteUser
	// todo GetUsers
}

func (s *Server) checkAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := &apiUser.CustomResponseWriter{ResponseWriter: w, StatusCode: 200}

		s.userHandler.CheckAccess(res, r)
		if res.StatusCode != 200 {
			apiJson.JSON(w, response.Error(res.Err, res.StatusCode))
			return
		}

		// todo read tokenData
		tokenData := r.Context().Value("tokenData")
		fmt.Println(tokenData)

		next.ServeHTTP(w, r)
	}
}

//func (s *Server) CheckAccess(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		// get token
//		tokenString := r.Header.Get("Authorization")
//		tokenData, err := s.CheckToken(tokenString)
//		if err != nil {
//			s.userHandler.GetAccessToken(w, r)
//			// нужна ошибка если внутри произойдт ошибка, или токен сгенерирован но не добавлен
//			// иначе без проверки вызывает
//		}
//
//		var ctx = context.WithValue(context.Background(), "tokenData", tokenData)
//		next.ServeHTTP(w, r.WithContext(ctx))
//	}
//}
//
//func (s *Server) CheckToken(tokenString string) (*apiUserModel.TokenData, error) {
//	token, err := utils.GetToken(tokenString)
//	if err != nil {
//		return nil, err
//	}
//
//	data, err := utils.VerifyToken(token, s.cfg.JwtSecret())
//	if err != nil {
//		return nil, err
//	}
//
//	return data, nil
//}

// todo найти место для этой хуйнии переписать ее

// func (s *Server) Authorization(next http.HandlerFunc) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var err error
//
//		defer func() {
//			if err != nil {
//				err := s.userHandler.GetAccessTokenOrError(w, r)
//				if err != nil {
//					return
//				}
//				tokenString, err := utils.GetToken(r)
//				if err != nil {
//					return
//				}
//				_, err = utils.VerifyToken(tokenString, s.cfg.JwtSecret())
//				if err != nil {
//					return
//					// apiJson.JSON(w, response.Error(err, http.StatusUnauthorized))
//				}
//
//				next.ServeHTTP(w, r)
//			}
//		}()
//
//		tokenString, err := utils.GetToken(r)
//		if err != nil {
//			return
//		}
//
//		_, err = utils.VerifyToken(tokenString, s.cfg.JwtSecret())
//		if err != nil {
//			return
//			// apiJson.JSON(w, response.Error(err, http.StatusUnauthorized))
//		}
//
//		next.ServeHTTP(w, r)
//	}
//}
