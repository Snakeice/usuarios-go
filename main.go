package main

import (
	r "github.com/snakeice/usuarios/routers"
	 "github.com/labstack/echo/middleware"
	"github.com/ikeikeikeike/pongor"
)

func main() {
	p := pongor.GetRenderer()
	p.Directory = "views"
	r.App.Renderer = p
	r.App.Use(middleware.Logger())
	r.App.Logger.Fatal(r.App.Start(":3000"))
}
