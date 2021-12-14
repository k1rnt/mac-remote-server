package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/valyala/fasthttp"

	model "github.com/k1rnt/mac-remote-server/model/request"
	"github.com/k1rnt/mac-remote-server/pkg/mac"
)

func Keycode() echo.HandlerFunc {
	return func(c echo.Context) error {
		var k model.KeycodeResquest
		if err := c.Bind(&k); err != nil {
			return c.String(fasthttp.StatusBadRequest, err.Error())
		}
		return c.String(fasthttp.StatusCreated, mac.Keycode(k.Code))
	}
}

func Keystroke() echo.HandlerFunc {
	return func(c echo.Context) error {
		var k model.KeyStrokeResquest
		if err := c.Bind(&k); err != nil {
			return c.String(fasthttp.StatusBadRequest, err.Error())
		}
		return c.String(fasthttp.StatusCreated, mac.Keystroke(k.IsSpecial, k.Key))
	}
}
