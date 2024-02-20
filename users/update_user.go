package users

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	if result := h.DB.First(&user, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	user.FirstName = body.FirstName
	user.LastName = body.LastName

	h.DB.Save(&user)

	ctx.JSON(http.StatusOK, &user)
}
