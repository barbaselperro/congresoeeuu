package main

import (
	"github.com/barbaselperro/Echo/rutas"
	"github.com/labstack/echo/v4"
)

func main() {
	route := echo.New()
	route.GET("/", rutas.Index)
	route.POST("/tutoken", rutas.Tutoken)
	route.POST("/acceso", rutas.Acceso)
	route.GET("/senadores", rutas.Tenadores)
	route.GET("/house", rutas.House)
	route.Start(":3000")
}
