package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"

	"github.com/k1rnt/mac-remote-server/api"
	"github.com/k1rnt/mac-remote-server/models"
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

func Keycode() echo.HandlerFunc {
	return func(c echo.Context) error {
		var k models.Keycode
		if err := c.Bind(&k); err != nil {
			return c.String(fasthttp.StatusBadRequest, err.Error())
		}
		return c.String(fasthttp.StatusCreated, api.Keycode(k.Code))
	}
}
