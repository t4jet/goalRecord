package router

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)
func SetRouter(e *echo.Echo) error {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} ${host} ${method} ${uri} ${status} ${header}\n",
		Output: os.Stdout,
	}))
  // ミドルウェアを設定
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // ルートを設定
	e.GET("/", hello)
	api := e.Group("/api")
		apiTasks := api.Group("/tasks")
			apiTasks.GET("", GetTaskHandler)
			apiTasks.POST("", AddTaskHandler)
			apiTasks.PUT("/:taskID", ChangeFinishedTaskHandler)
			apiTasks.DELETE("/:taskID", DeleteTaskHandler)

  // サーバーをポート番号1323で起動
	err := e.Start(":8080")
	return err
}
// ハンドラーを定義
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}
