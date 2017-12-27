package routers

import (
	"github.com/labstack/echo"
	"github.com/snakeice/usuarios/controllers"
)

var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/atualizar/:id", controllers.Update)

	api := App.Group("/v1")

	api.POST("/insert", controllers.Inserir)
	api.DELETE("/delete/:id", controllers.Deletar)
	api.PUT("/update/:id", controllers.Atualizar)

	}

