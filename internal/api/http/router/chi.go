package router

import (
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type router struct {
	r chi.Router
}

func New() Router {
	cr := &router{
		r: chi.NewRouter(),
	}

	return cr
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.r.ServeHTTP(w, req)
}

func (r *router) AddMiddleware(middleware func(handler http.Handler) http.Handler) {
	r.r.Use(middleware)
}

func (r *router) AddRoute(method, route string, handler http.HandlerFunc) {
	switch method {
	case http.MethodGet:
		r.r.Get(route, handler)
	case http.MethodPost:
		r.r.Post(route, handler)
	case http.MethodDelete:
		r.r.Delete(route, handler)
		// todo other
	}

}

func (r *router) AddGroup(groupName string, handler map[string]map[string]http.HandlerFunc) {
	r.r.Route(groupName, func(rc chi.Router) {
		for k, v := range handler {
			switch k {
			case http.MethodGet:
				for k, v := range v {
					r.r.Get(k, v)
				}
			case http.MethodPost:
				for k, v := range v {
					r.r.Post(k, v)
				}
			case http.MethodDelete:
				for k, v := range v {
					r.r.Delete(k, v)
				}
			}
		}
	})

}

// todo delete
func (r *router) CurrentRoute(ctx context.Context) string {
	return chi.RouteContext(ctx).RoutePattern()
}
