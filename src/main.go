package main

import (
	"../src/router"
	"fmt"
)

func main() {
	echo := router.New()
	if err := echo.Start(":8000"); err != nil {
		fmt.Println("Error")
	}
}
