package books

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) getBooks(ctx *gin.Context) {
	var books []models.Book

	author := ctx.Query("author")
	title := ctx.Query("title")

	if result := h.DB.Where("author LIKE ?", "%"+author+"%").Where("title LIKE ?", "%"+title+"%").Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}
