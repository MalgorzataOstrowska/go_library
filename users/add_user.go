package users

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type addUserRequestBody struct {
	firstName string
	lastName  string
}

func (h handler) addUser(ctx *gin.Context) {
	body := addUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.FirstName = body.firstName
	user.LastName = body.lastName

	if result := h.DB.Create(&user); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &user)
}
