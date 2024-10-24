package books

import (
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"github.com/gin-gonic/gin"
)

func (h DBHandler) GetBooks(ctx *gin.Context){
 	var books []model.Book
	if result := h.db.Find(&books); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &books)
}