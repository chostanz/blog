package main

import (
	"blog/routes"

	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := routes.Route()

	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":8080"))

}
