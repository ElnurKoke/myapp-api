package handler

import (
	"elestial/config"
	"elestial/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Router  *gin.Engine
	Service *service.Service
	Cfg     *config.Config
}

func NewHandler(services *service.Service) *Handler {
	router := gin.Default()
	return &Handler{
		Router:  router,
		Service: services,
	}
}

func (h *Handler) InitRoutes() http.Handler {
	h.Router.POST("/register", h.register)
	h.Router.POST("/login", h.login)
	h.Router.POST("/logout", h.logout)
	h.Router.POST("/refresh", h.refresh)
	//h.Router.GET("/home", h.AuthMiddleware(), h.home)

	return h.Router
}
