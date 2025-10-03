package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"piscord-backend/config"
	"piscord-backend/services"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	mongoService := services.NewMongoService(cfg.MongoURI)

	if err := mongoService.Connect(); err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}
	defer func() {
		if err := mongoService.Disconnect(); err != nil {
			log.Println("Error disconnecting from MongoDB: ", err)
		}
	}()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	srv := &http.Server{
		Addr:    "0.0.0.0:" + cfg.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}
