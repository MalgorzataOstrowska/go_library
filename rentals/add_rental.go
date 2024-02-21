package rentals

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type addRentalRequestBody struct {
	userId int
	bookId int
}

func (h handler) addRental(ctx *gin.Context) {
	body := addRentalRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var rental models.Rental

	rental.UserId = body.userId
	rental.BookId = body.bookId
	rental.RentalDate = time.Now()
	rental.ReturnData = time.Now().AddDate(0, 0, 14)

	if result := h.DB.Create(&rental); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &rental)
}
