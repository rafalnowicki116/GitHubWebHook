package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetLogger(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}][remote_ip:${remote_ip}][status: ${status}][method: ${method}][url: ${host}${path}]` + "\n",
	}))
}
