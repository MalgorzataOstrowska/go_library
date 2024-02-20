package rentals

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type UpdateRentalRequestBody struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}

func (h handler) UpdateRental(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateRentalRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var rental models.Rental

	if result := h.DB.First(&rental, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	rental.UserId = body.UserId
	rental.BookId = body.BookId
	rental.RentalDate = time.Now()
	rental.ReturnData = time.Now().AddDate(0, 0, 14)

	h.DB.Save(&rental)

	ctx.JSON(http.StatusOK, &rental)
}
