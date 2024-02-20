package users

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type AddUserRequestBody struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (h handler) AddUser(ctx *gin.Context) {
	body := AddUserRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user models.User

	user.FirstName = body.FirstName
	user.LastName = body.LastName

	if result := h.DB.Create(&user); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &user)
}
