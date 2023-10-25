package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func SetRouter(e *echo.Echo) error {
  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルートを設定
  e.GET("/", hello)

  // サーバーをポート番号1323で起動
	err := e.Start(":8080")
	return err
}
// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
