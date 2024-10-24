package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Mohammedbilal1998/go-postgres/pkg/books"
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/config"
	"github.com/Mohammedbilal1998/go-postgres/pkg/common/db"
	"github.com/gin-gonic/gin"
)

func main () {
	//read config
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("unable to read the config")
	}
	// get handler for gin apis
	db, err := db.Init(config.DBUrl)
	if err != nil {
		log.Fatal(err)
	}

	// prepare router
	router := gin.Default()
	books.RegesterRoutes(router, db)

	//health chek
	router.GET("/", func(ctx *gin.Context){
		ctx.JSON(http.StatusOK, gin.H{
			"port": config.Port,
			"DBUrl": config.DBUrl,
		})
	})

	log.Println(config)
	srv := &http.Server{
		Addr: config.Port,
		Handler: router,
	}

	// graceful shutdown
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt,os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func(){
		<-interupt
		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal(err)
		}
		fmt.Println("gracefully shutting down")
	}()
	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}