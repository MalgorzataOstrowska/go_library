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
	routes.POST("/", h.AddRental)
	routes.GET("/", h.GetRentals)
	routes.GET("/:id", h.GetRental)
	routes.PUT("/:id", h.UpdateRental)
	routes.DELETE("/:id", h.DeleteRental)
}
