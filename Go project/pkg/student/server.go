package student

import (
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/model"
	"github.com/gin-gonic/gin"
)

type CreateStudentRequest struct {
	ID uint  `json:"id"`
	Name string  `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}


func (h DBHandler) GetAllStudents(ctx *gin.Context){
	students := []model.Student{}
	if result := h.db.Find(&students); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &students)
}

func (h DBHandler) CreateStudent(ctx *gin.Context) {
	var student  CreateStudentRequest
	err := ctx.BindJSON(&student)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	if result := h.db.Create(&student); result.Error != nil {
		ctx.AbortWithError(http.StatusNotModified, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &student)
}