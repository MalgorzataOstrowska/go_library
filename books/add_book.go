package books

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type addBookRequestBody struct {
	title         string
	author        string
	description   string
	publication   time.Time
	numberOfPages int
}

func (h handler) addBook(ctx *gin.Context) {
	body := addBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	book.Title = body.title
	book.Author = body.author
	book.Description = body.description
	book.Publication = body.publication
	book.NumberOfPages = body.numberOfPages

	if result := h.DB.Create(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &book)
}
