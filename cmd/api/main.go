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
	"github.com/TroJanBoi/assembly-visual-backend/internal/model"
	"github.com/TroJanBoi/assembly-visual-backend/internal/scheduler"
	"github.com/TroJanBoi/assembly-visual-backend/internal/server"
	"github.com/robfig/cron/v3"
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

func startScheduler(logger *zap.Logger) {
	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Bangkok", 7*3600)))
	_, err := c.AddFunc("* * * * *", func() {
		logger.Info("Scheduler running...")

		scheduler.CleanupSoftDeletedUsers(&model.User{}, 3)         // Clean up users soft-deleted more than 3 days ago
		scheduler.CleanupExpiredInvitations(&model.Invitation{}, 1) // Clean up invitations older than 3 days
	})

	if err != nil {
		logger.Error("Failed to start scheduler: " + err.Error())
		return
	}

	c.Start()
	logger.Info("Scheduler started")
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

	startScheduler(zap)

	srv := server.NewServer()
	go gracefully(srv, zap, gracefulShutdownDuration)

	port := strconv.Itoa(config.PORT)
	zap.Info("Starting server on port " + port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}

	log.Println("server exited gracefully")
}
