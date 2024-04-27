package router

import (
	"context"
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	AddRoute(method, route string, handler http.HandlerFunc)
	AddMiddleware(middleware func(handler http.Handler) http.Handler)
	AddGroup(groupName string, handler map[string]map[string]http.HandlerFunc)
	CurrentRoute(ctx context.Context) string
}
