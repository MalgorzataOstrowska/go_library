package rentals

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetRentals(ctx *gin.Context) {
	var rentals []models.Rental

	if result := h.DB.Preload("User").Preload("Book").Find(&rentals); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &rentals)
}
