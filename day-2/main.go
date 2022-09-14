package main

import (
	"github.com/mahendrabp/alterra-agmc/day-2/config"
	"github.com/mahendrabp/alterra-agmc/day-2/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
