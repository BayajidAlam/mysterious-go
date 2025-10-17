package cmd

import (
	"go.mod/config"
	"go.mod/rest"
)

func Serve() {
	cnf := config.GetConfig()
	rest.Start(cnf)
}
