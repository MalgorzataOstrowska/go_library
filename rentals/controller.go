package rentals

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

	routes := router.Group("/rentals")
	routes.POST("/", h.addRental)
	routes.GET("/", h.getRentals)
	routes.GET("/:id", h.getRental)
	routes.PUT("/:id", h.updateRental)
	routes.DELETE("/:id", h.deleteRental)
}
