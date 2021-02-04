package api

import (
	"../api/handlers"
	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.POST("/webhook", handlers.WebHook)
}
