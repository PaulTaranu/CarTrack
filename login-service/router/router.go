package router

import (
	"github.com/PaulTaranu/CarTrack/login-service/handlers"
	"github.com/labstack/echo/v4"
)

func InitRoutes() *echo.Echo {
	e := echo.New()

	e.POST("/register", handlers.RegisterHandler)
	e.POST("/login", handlers.LoginHandler)

	return e
}
