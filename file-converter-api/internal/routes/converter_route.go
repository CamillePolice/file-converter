package routes

import (
	"file-converter-api/internal/server"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ConverterRoutes struct {
	server *server.Server
}

func NewConverterRoutes(s *server.Server) *ConverterRoutes {
	return &ConverterRoutes{server: s}
}

func (r *ConverterRoutes) Register(v1 *echo.Group) {
	fileConverters := v1.Group("/file/converter")

	fileConverters.GET("", r.ConvertFile)
	//users.GET("", r.handleGetUsers)
	//users.GET("/:id", r.handleGetUser)
	//users.POST("", r.handleCreateUser)
	//users.PUT("/:id", r.handleUpdateUser)
	//users.DELETE("/:id", r.handleDeleteUser)
}

func (r *ConverterRoutes) ConvertFile(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Convert file",
	})
}
