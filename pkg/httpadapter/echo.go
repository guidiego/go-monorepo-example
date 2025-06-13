package httpadapter

import "github.com/labstack/echo/v4"

func EchoAdapter(controller IHttpAdapter) echo.HandlerFunc {
	return func(c echo.Context) error {
		return controller.Handle(c.Response().Writer, c.Request())
	}
}
