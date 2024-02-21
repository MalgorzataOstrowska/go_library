package books

import (
	"net/http"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

func (h handler) deleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&book)

	ctx.Status(http.StatusOK)
}
