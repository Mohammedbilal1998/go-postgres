package books

import (
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"github.com/gin-gonic/gin"
)


func (h DBHandler) DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	var book model.Book

	if result := h.db.First(&book, bookId); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := h.db.Delete(&book); result.Error != nil {
		ctx.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &book)
}