package books

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DBHandler struct {
	db *gorm.DB

}
func RegesterRoutes(router *gin.Engine, handler *gorm.DB) {
	h := DBHandler{db: handler}
	
	routes := router.Group("/books")
	routes.POST("/", h.AddBook)
	routes.GET("/:id", h.GetBook)
	routes.GET("/", h.GetBooks)
	routes.PUT("/:id", h.UpdateBook)
	routes.DELETE("/:id", h.DeleteBook)
}
