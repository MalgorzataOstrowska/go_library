package users

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type updateUserRequestBody struct {
	firstName string
	lastName  string
}

func (h handler) updateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	body := updateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	user.FirstName = body.firstName
	user.LastName = body.lastName

	h.DB.Save(&user)

	ctx.JSON(http.StatusOK, &user)
}
