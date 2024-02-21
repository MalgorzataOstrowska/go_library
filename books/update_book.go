package books

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type updateBookRequestBody struct {
	title         string
	author        string
	description   string
	publication   time.Time
	numberOfPages int
}

func (h handler) updateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := updateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.title
	book.Author = body.author
	book.Description = body.description
	book.Publication = body.publication
	book.NumberOfPages = body.numberOfPages

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}
