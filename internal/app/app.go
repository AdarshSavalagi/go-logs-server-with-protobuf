package app

import (
	"fmt"

	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/config"
	kafka_mem "github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/kafka"
	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/logger"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type App struct {
	Config      *config.Config
	Logger      *logrus.Logger
	Router      *gin.Engine
	KafkaWriter map[string]*kafka.Writer
}

func InitApp() (*App, error) {
	logger := logger.InitLogger()

	// Initialize config
	config, err := config.InitConfig()
	if err != nil {
		logger.Fatalf("Failed to initialize config: %v", err) // Fix error formatting
		return nil, err
	}

	// Initialize Gin router
	router := gin.Default()

	// Initialize Kafka writers
	kafkaWriter, err := kafka_mem.InitKafkaWriters(&config.Kafka)

	// Check if Kafka writers were initialized successfully
	if err != nil {
		logger.Fatalf("Failed to initialize Kafka writers: %v", err) // Fix error formatting
		return nil, err
	}

	// Return the app instance
	app := &App{
		Config:      config,
		Logger:      logger,
		Router:      router,
		KafkaWriter: kafkaWriter,
	}
	SetupRoutes(app)

	return app, nil
}

func (app *App) Run() {
	port := app.Config.Server.Port
	if port == 0 {
		port = 8000
	}
	address := fmt.Sprintf(":%d", port)

	// Start the HTTP server.
	if err := app.Router.Run(address); err != nil {
		app.Logger.Fatalf("Failed to start server: %v", err)
	}
}
