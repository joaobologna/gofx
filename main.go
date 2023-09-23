package main

import (
	"github.com/joaobologna/gofx/log"
	"github.com/joaobologna/gofx/server"
	"github.com/joaobologna/gofx/ucs"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		log.Module,
		server.Module,
		ucs.Module,
	).Run()
}
