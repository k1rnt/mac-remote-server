package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/k1rnt/mac-remote-server/handler"
)

func Init() *echo.Echo {
	server := echo.New()

	// middleware
	server.Use(middleware.Logger())
	server.Use(middleware.Gzip())
	server.Use(middleware.Recover())
	server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	// Routes
	api := server.Group("/api")
	{
		api.POST("/unsleep", handler.UnSleep())
		api.POST("/sleep", handler.Sleep())
		api.POST("/keycode", handler.Keycode())
		api.POST("/keystroke", handler.Keystroke())
	}

	return server

}
