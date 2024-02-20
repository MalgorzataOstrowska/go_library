package users

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUserRentals(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	if result := h.DB.Preload("Rentals").First(&user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &user)
}
