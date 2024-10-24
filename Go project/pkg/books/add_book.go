package books

import (
	"fmt"
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"github.com/gin-gonic/gin"
)

type CreateBookRequestBody struct {
    Title string `json:"title"` 
    Author string `json:"author"`
    Description string `json:"description"`
}

func (h DBHandler) AddBook(ctx *gin.Context){
    body := CreateBookRequestBody{}
    fmt.Println("CALLED POST")

    if err := ctx.BindJSON(&body); err != nil {
        fmt.Println("invalid request details")
    }
    fmt.Println(body)
    book := model.Book{}
    book.Title = body.Title
    book.Author = body.Author
    book.Description = body.Description

    if result := h.db.Create(&book); result.Error != nil {
        ctx.AbortWithError(http.StatusNotImplemented, result.Error)
        return
    }
    ctx.JSON(http.StatusCreated, &body)
}