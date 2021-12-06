package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"

	mac "github.com/k1rnt/mac-remote-server/pkg/macAPI"
)

func UnSleep() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.String(fasthttp.StatusCreated, mac.UnSleep())
	}
}

func Sleep() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.String(fasthttp.StatusCreated, mac.Sleep())
	}
}

func Lock() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := "12 using {command down, control down}"
		return c.String(fasthttp.StatusCreated, mac.Keycode(code))
	}
}
