package routers

import (
	"github.com/gin-gonic/gin"
	"shortener/internal/service"
)

type Handler struct {
	services *service.Services
}

func AppHandler(services *service.Services) *Handler {
	return &Handler{
		services: services}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()
	router.GET("api/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "pong"})
	})
	return router
}
