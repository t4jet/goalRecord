package main

import (
	"github.com/labstack/echo/v4"
	"main/router"
)

func main() {
  // インスタンスを作成
  e := echo.New()
	router.SetRouter(e)
}
