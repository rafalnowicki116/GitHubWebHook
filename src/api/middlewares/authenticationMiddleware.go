package middlewares

import (
	"fmt"
	"github.com/labstack/echo"
)

func SetAuthenticationMiddleware(group *echo.Echo) {
	fmt.Println("Authentication middleware")
}
