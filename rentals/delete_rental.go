package rentals

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) DeleteRental(ctx *gin.Context) {
	id := ctx.Param("id")

	var rental models.Rental

	if result := h.DB.First(&rental, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&rental)

	ctx.Status(http.StatusOK)
}
