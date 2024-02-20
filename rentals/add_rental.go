package rentals

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type AddRentalRequestBody struct {
	UserId int `json:"user_id"`
	BookId int `json:"book_id"`
}

func (h handler) AddRental(ctx *gin.Context) {
	body := AddRentalRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var rental models.Rental

	rental.UserId = body.UserId
	rental.BookId = body.BookId
	rental.RentalDate = time.Now()
	rental.ReturnData = time.Now().AddDate(0, 0, 14)

	if result := h.DB.Create(&rental); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &rental)
}
