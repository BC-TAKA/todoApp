package handler

import "github.com/todoApp/api/domain/service"

type Handler interface {
	TODOHandler
}

type handler struct {
	TODOHandler
}

// value実態無し
func NewHandler(value service.TODOService) Handler {
	return &handler{
		TODOHandler: NewTODOHandler(value),
	}
}
