package logs_api

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

func Routes(r *gin.Engine, kafkaWriters map[string]*kafka.Writer, logger *logrus.Logger) {
	logHandler := NewLogHandler(kafkaWriters,logger)

	router := r.Group("/api/v1/logs")
	router.POST("/upload-logs/", logHandler.UploadLogs)
	router.POST("/upload-events/", logHandler.UploadEvents)
	router.POST("/upload-context/", logHandler.UploadContext)
}
