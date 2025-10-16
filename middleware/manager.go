package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

// ServeHTTP implements http.Handler.
func (m Middleware) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

type Manager struct {
	globalMiddleware []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddleware: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globalMiddleware = append(mngr.globalMiddleware, middlewares...,
	)
}

func (mngr *Manager) With(handler http.Handler, middlewares ...Middleware) http.Handler {

	h := handler
	for _, middleware := range middlewares {
		h = middleware(h)
	}

	for _, globalMiddleware := range mngr.globalMiddleware {
		h = globalMiddleware(h)
	}
	return h
}

func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {

	h := handler
	for _, middleware := range mngr.globalMiddleware {
		h = middleware(h)
	}

	return h
}
