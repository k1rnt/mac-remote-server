package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"

	"github.com/k1rnt/mac-remote-server/api"
)

func UnSleep() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.String(fasthttp.StatusCreated, api.UnSleep())
	}
}

func Sleep() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.String(fasthttp.StatusCreated, api.Sleep())
	}
}
