package main

import (
	"github.com/mahendrabp/alterra-agmc/day-3/config"
	middleware "github.com/mahendrabp/alterra-agmc/day-3/middlewares"
	"github.com/mahendrabp/alterra-agmc/day-3/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	middleware.LogMiddleware(e)
	e.Logger.Fatal(e.Start(":8000"))
}
