package routes

import (
	"file-converter-api/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Routes struct {
	server              *server.Server
	fileConverterRoutes *ConverterRoutes
}

func NewRoutes(s *server.Server) *Routes {
	return &Routes{
		server:              s,
		fileConverterRoutes: NewConverterRoutes(s),
	}
}

func (r *Routes) Setup(e *echo.Echo) {
	// API v1 group
	v1 := e.Group("/api/v1")

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/api/v1")
	})

	v1.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "it works",
		})
	})

	// Register all routes
	r.fileConverterRoutes.Register(v1)
}
