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
	routes.POST("/", h.AddUser)
	routes.GET("/", h.GetUsers)
	routes.GET("/:id", h.GetUser)
	routes.GET("/:id/rentals", h.GetUserRentals)
	routes.PUT("/:id", h.UpdateUser)
	routes.DELETE("/:id", h.DeleteUser)
}
