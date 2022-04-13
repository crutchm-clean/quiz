package handler

import (
	"Kahoot/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/reg", h.signUp)
		auth.POST("/login", h.signIn)
	}
	test := router.Group("/test", h.userIdentity)
	{
		test.GET("/", h.GetTests)
		test.POST("/create", h.CreateTest)
	}
	return router
}
