package main

import (
	"context"
	"go-payment-srv/handler"
	"go-payment-srv/repository"
	"go-payment-srv/usecase"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	route := gin.Default()

	// init logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Panic("failed to initialize logger", err)
	}
	sugared := logger.Sugar()

	// init gorm
	dsn := "host=db user=transaction_user password=password dbname=transaction port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	g, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to initialize db", err)
	}

	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	repo := repository.NewRepo(g)
	uc := usecase.NewUsecase(repo, sugared)
	h := handler.NewHandler(uc, "http://node-account-srv:8000")

	auth := func() gin.HandlerFunc {
		return h.Auth
	}

	payments := route.Group("/payments")
	{
		payments.Use(auth())
		payments.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"ping": "pong",
			})
		})
		payments.POST("/send", h.Send)
		payments.POST("/withdraw", h.Withdraw)
	}

	srv := &http.Server{
		Addr:    ":8001",
		Handler: route.Handler(),
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
