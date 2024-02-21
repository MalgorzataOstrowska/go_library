package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := router.Group("/books")
	routes.POST("/", h.addBook)
	routes.GET("/", h.getBooks)
	routes.GET("/:id", h.getBook)
	routes.PUT("/:id", h.updateBook)
	routes.DELETE("/:id", h.deleteBook)
}
