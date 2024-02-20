package rentals

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetRental(ctx *gin.Context) {
	id := ctx.Param("id")

	var rental models.Rental

	if result := h.DB.Preload("User").Preload("Book").First(&rental, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &rental)
}
