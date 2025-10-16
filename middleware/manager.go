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

	n := handler
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, globalMiddleware := range mngr.globalMiddleware {
		n = globalMiddleware(n)
	}
	return n

}
