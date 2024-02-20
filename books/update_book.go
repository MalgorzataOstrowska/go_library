package books

import (
	"net/http"
	"time"

	"library/common/models"

	"github.com/gin-gonic/gin"
)

type UpdateBookRequestBody struct {
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Description   string    `json:"description"`
	Publication   time.Time `json:"publication"`
	NumberOfPages int       `json:"number_of_pages"`
}

func (h handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description
	book.Publication = body.Publication
	book.NumberOfPages = body.NumberOfPages

	h.DB.Save(&book)

	ctx.JSON(http.StatusOK, &book)
}
