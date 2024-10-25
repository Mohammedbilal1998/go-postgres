package student

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBHandler struct{
	db *gorm.DB
}

func SetRoutes(route *gin.Engine, client *gorm.DB){
	h := DBHandler{
		db: client,
	}
	router := route.Group("student")

	router.GET("/students", h.GetAllStudents)
	router.POST("/student", h.CreateStudent)
}
