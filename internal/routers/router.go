package routers

import (
	"net/http"
	"shortener/internal/service"
)

type Handler struct {
	services *service.Services
}

func AppHandler(services *service.Services) *Handler {
	return &Handler{
		services: services}
}

func (h *Handler) Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.HandlePostShorten)
	return mux
}
