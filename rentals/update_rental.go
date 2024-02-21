package rentals

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type updateRentalRequestBody struct {
	userId int
	bookId int
}

func (h handler) updateRental(ctx *gin.Context) {
	id := ctx.Param("id")
	body := updateRentalRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var rental models.Rental

	if result := h.DB.First(&rental, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	rental.UserId = body.userId
	rental.BookId = body.bookId
	rental.RentalDate = time.Now()
	rental.ReturnData = time.Now().AddDate(0, 0, 14)

	h.DB.Save(&rental)

	ctx.JSON(http.StatusOK, &rental)
}
