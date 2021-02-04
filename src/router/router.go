package router

import (
	"../api"
	"../api/middlewares"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	middlewares.SetAuthenticationMiddleware(e)
	middlewares.SetLogger(e)
	api.MainGroup(e)

	return e
}
