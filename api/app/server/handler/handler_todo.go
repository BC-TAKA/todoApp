package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/todoApp/api/domain/service"
)

type TODOHandler interface {
	// 関数を宣言
	TODOFunc(c echo.Context) error
}

type todoHandler struct {
	todo service.TODOService
}

// valueの実態無し
func NewTODOHandler(value service.TODOService) TODOHandler {
	return &todoHandler{
		todo: value,
	}
}

func (td *todoHandler) TODOFunc(c echo.Context) error {
	// service層からの結果をrespに格納して返す
	return c.JSON(http.StatusOK, resp)
}
