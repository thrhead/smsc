package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	
	"smsc/internal/api"
	"smsc/internal/config"
	"smsc/internal/db"
	"smsc/internal/protocols/smpp"
	"smsc/internal/protocols/sigtran"
	"smsc/internal/services/monitoring"
	"smsc/internal/services/queue"
	"smsc/internal/services/routing"
	"smsc/pkg/logger"
)

var (
	configFile string
	log        *logrus.Logger
)

func init() {
	flag.StringVar(&configFile, "config", "config/config.yaml", "path to configuration file")
}

func main() {
	flag.Parse()

	// Load configuration
	cfg, err := config.Load(configFile)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// Initialize logger
	log, err = logger.New(cfg.Logging)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// Set up context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize database
	database, err := db.New(cfg.Database, log)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Initialize database schema
	if err := database.InitSchema(ctx); err != nil {
		log.Fatalf("Failed to initialize database schema: %v", err)
	}

	// Initialize services
	monitoringService := monitoring.New(cfg.Monitoring, log)
	if err := monitoringService.Start(ctx); err != nil {
		log.Fatalf("Failed to start monitoring service: %v", err)
	}

	queueService := queue.New(cfg.Queue, log)
	if err := queueService.Start(ctx); err != nil {
		log.Fatalf("Failed to start queue service: %v", err)
	}

	routingService := routing.New(cfg.Routing, log)
	if err := routingService.Start(ctx); err != nil {
		log.Fatalf("Failed to start routing service: %v", err)
	}

	// Initialize protocol handlers
	smppServer := smpp.New(cfg.SMPP, log)
	if err := smppServer.Start(); err != nil {
		log.Fatalf("Failed to start SMPP server: %v", err)
	}

	sigtranStack := sigtran.New(cfg.Sigtran, log)
	if err := sigtranStack.Start(); err != nil {
		log.Fatalf("Failed to start Sigtran stack: %v", err)
	}

	// Initialize API server
	apiServer := api.New(api.Config{
		Host:           cfg.Server.Host,
		Port:           cfg.Server.Port,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}, log)

	if err := apiServer.Start(); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

	// Set up graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	sig := <-sigChan
	log.Infof("Received signal %v, initiating shutdown", sig)

	// Create shutdown context with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Trigger graceful shutdown
	if err := apiServer.Stop(shutdownCtx); err != nil {
		log.Errorf("API server shutdown error: %v", err)
	}

	if err := smppServer.Stop(); err != nil {
		log.Errorf("SMPP server shutdown error: %v", err)
	}

	if err := sigtranStack.Stop(); err != nil {
		log.Errorf("Sigtran stack shutdown error: %v", err)
	}

	// Cancel context to stop all services
	cancel()

	log.Info("Shutdown complete")
} 