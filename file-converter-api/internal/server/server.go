package server

import (
	"file-converter-api/internal/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo   *echo.Echo
	config *config.ServerConf
}

func NewServer(config *config.ServerConf) *Server {
	echoServer := echo.New()

	// Configure middleware
	echoServer.Use(middleware.Logger())
	echoServer.Use(middleware.Recover())

	return &Server{
		echo:   echoServer,
		config: config,
	}
}

func (s *Server) Start() error {
	return s.echo.Start(s.config.Port)
}

func (s *Server) Echo() *echo.Echo {
	return s.echo
}
