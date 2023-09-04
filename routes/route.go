package routes

import (
	"ukzuhdi/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/product", controllers.GetProductController)
	e.POST("/product", controllers.AddProductController)
	e.DELETE("/product/:id", controllers.DeleteProductController)
	e.PUT("/product/:id", controllers.UpdateProductController)
}
