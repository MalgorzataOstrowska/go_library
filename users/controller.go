package users

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

	routes := router.Group("/users")
	routes.POST("/", h.addUser)
	routes.GET("/", h.getUsers)
	routes.GET("/:id", h.getUser)
	routes.GET("/:id/rentals", h.getUserRentals)
	routes.PUT("/:id", h.updateUser)
	routes.DELETE("/:id", h.deleteUser)
}
