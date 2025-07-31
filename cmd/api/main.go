package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/TroJanBoi/assembly-visual-backend/internal/conf"
	"github.com/TroJanBoi/assembly-visual-backend/internal/server"
	"go.uber.org/zap"
)

const (
	gracefulShutdownDuration = 10 * time.Second
)

func gracefully(srv *http.Server, log *zap.Logger, shutdownTimeout time.Duration) {
	{
		ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer cancel()
		<-ctx.Done()
	}

	log.Info("Shutting down server gracefully...")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Info("HTTP server shutdown: " + err.Error())
	}
}

// @title Assembly Visual API documentation
// @version 1.0
// @description This is the API documentation for the Assembly Visual project.
// @termsOfService http://swagger.io/terms/
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config := conf.NewConfig()
	zap, _ := zap.NewProduction()
	defer zap.Sync()
	srv := server.NewServer()

	go gracefully(srv, zap, gracefulShutdownDuration)

	port := strconv.Itoa(config.PORT)
	zap.Info("Starting server on port " + port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("server exited gracefully")
}
