package books

import (
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"github.com/gin-gonic/gin"
)

func (h DBHandler) UpdateBook(ctx *gin.Context){
	bookId := ctx.Param("id")
	var body model.Book
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var book model.Book
	if result := h.db.First(&book, bookId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	book.Title = body.Title
	book.Author = body.Author
	book.Description = body.Description

	if result := h.db.Save(book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotModified, result.Error)
	}

	ctx.JSON(http.StatusOK, book)
}