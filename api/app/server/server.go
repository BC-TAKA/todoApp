package server

import (
	"github.com/labstack/echo/v4"
	"github.com/todoApp/api/app/server/handler"
)

type Server struct {
	echo *echo.Echo
}

func New(e *echo.Echo) *Server {
	return &Server{
		echo: e,
	}
}

func (s *Server) Handlers(h handler.Handler) *Server {
	s.echo.GET("api/test", h.TODOFunc)
	return s
}

// ポート番号の設定を返す関数
func (s *Server) Start() error {
	port := "1234"
	return s.echo.Start(":" + port)
}
