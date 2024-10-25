package main

import (
	// "net/http"

	"fmt"
	"log"
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/config"
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/db"
	"github.com/Mohammedbilal1998/go-postgres/pkg/student"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.ReadConfig()
	fmt.Println("config: ",config)

	route := gin.Default()

	route.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "health check successful")
	})
	dbClient, err := db.GetClient(config)
	if err != nil {
		log.Fatal(err)
	}

	student.SetRoutes(route, dbClient)

	route.Run()




}