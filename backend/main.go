package main

import (
	"github.com/labstack/echo/v4"
	"main/router"
	"main/model"
)

func main() {
	sqlDB := model.DBConnection()
	defer sqlDB.Close()
  // インスタンスを作成
  e := echo.New()
	router.SetRouter(e)
}
