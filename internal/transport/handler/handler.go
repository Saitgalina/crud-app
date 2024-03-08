package handler

import (
	"github.com/Saitgalina/crud-app/internal/core/interface/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Sevice
}

func NewHandler(services *service.Sevice) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sing-up", h.singUp)
		auth.POST("/sing-in", h.singIn)
	}
	api := router.Group("api", h.userIdentity)
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getAllBooks)
			books.GET("/filter", h.getFilterBooks)
			books.GET("/sort", h.getSortBooks)
			//books.GET("/name", h.getBookByName)
			//books.GET("/:id", h.getBookById) //рабочий
			//вызывают панику потому что "/:id" в этом месте можно писать только :id
			//books.GET("/:name", h.getBookByName)
			//books.GET("/:author", h.getBookByAuthor)
		}
	}
	return router
}
