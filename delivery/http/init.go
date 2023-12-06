package http

import (
	"github.com/julienschmidt/httprouter"
)

type HttpHandler struct {
	router  *httprouter.Router
	usecase Usecases
}

func NewHTTP(router *httprouter.Router, usecase Usecases) *HttpHandler {
	return &HttpHandler{
		router:  router,
		usecase: usecase,
	}
}

func (h *HttpHandler) SetEndpoint() {
	h.router.POST("/v1/set", h.HandlerSETPM)
}
