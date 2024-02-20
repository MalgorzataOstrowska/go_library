package users

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(ctx *gin.Context) {
	var users []models.User

	first_name := ctx.Query("first_name")
	last_name := ctx.Query("last_name")

	if result := h.DB.Where("first_name LIKE ?", "%"+first_name+"%").Where("last_name LIKE ?", "%"+last_name+"%").Find(&users); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &users)
}
