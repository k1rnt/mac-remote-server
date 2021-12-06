package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"

	"github.com/k1rnt/mac-remote-server/api"
	model "github.com/k1rnt/mac-remote-server/model/request"
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

func Lock() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := "12 using {command down, control down}"
		return c.String(fasthttp.StatusCreated, api.Keycode(code))
	}
}

func Keycode() echo.HandlerFunc {
	return func(c echo.Context) error {
		var k model.KeycodeResquest
		if err := c.Bind(&k); err != nil {
			return c.String(fasthttp.StatusBadRequest, err.Error())
		}
		return c.String(fasthttp.StatusCreated, api.Keycode(k.Code))
	}
}

func Keystroke() echo.HandlerFunc {
	return func(c echo.Context) error {
		var k model.KeyStrokeResquest
		if err := c.Bind(&k); err != nil {
			return c.String(fasthttp.StatusBadRequest, err.Error())
		}
		return c.String(fasthttp.StatusCreated, api.Keystroke(k.IsSpecial, k.Key))
	}
}
