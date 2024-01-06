package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/todoApp/api/app/server"
	"github.com/todoApp/api/app/server/handler"
	"github.com/todoApp/api/domain/service"
	"github.com/todoApp/api/infra"
)

func NewHandler() handler.Handler {

	repo := infra.NewTODO()
	service := service.NewTODOService(repo)
	// repo := infra.New(&infra.HTTPClient{}, db)
	// service := service.New(repo)
	// 各層の構造体を渡す
	h := handler.NewHandler(service)

	return h
}

func main() {
	h := NewHandler()
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	server := server.New(e).Handlers(h)
	log.Fatal(server.Start())
}
