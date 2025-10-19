package middleware

import "go.mod/config"

type Middlewares struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}
