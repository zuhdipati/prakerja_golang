package main

import (
	"ukzuhdi/configs"
	"ukzuhdi/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	e.Start(":8080")
}
