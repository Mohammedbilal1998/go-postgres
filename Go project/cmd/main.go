package main

import (
	"log"
	"net/http"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/config"
	"github.com/gin-gonic/gin"
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/db"
	"github.com/Mohammedbilal1998/go-postgres/pkg/books"
)

func main () {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("unable to read the config")
	}
	router := gin.Default()
	dbHandler, err := db.Init(config.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	books.RegesterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"port": config.Port,
			"DBUrl": config.DBUrl,
		})
	})

	log.Println(config)
	router.Run(config.Port)
}