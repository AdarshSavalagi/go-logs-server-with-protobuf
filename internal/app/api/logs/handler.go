package logs_api

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"

	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type LogHandler struct {
	kafkaWriters map[string]*kafka.Writer
	logger       *logrus.Logger
}

func NewLogHandler(kafkaWriters map[string]*kafka.Writer, logger *logrus.Logger) *LogHandler {
	return &LogHandler{
		kafkaWriters: kafkaWriters,
		logger:       logger,
	}
}

func (h *LogHandler) UploadLogs(c *gin.Context) {
	writer, ok := h.kafkaWriters["logs"]
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Kafka writer not available"})
		return
	}
	// Debug: Kafka writer info
	h.logger.Info("Uploading logs: Kafka writer available")

	var reader io.ReadCloser
	var err error

	if c.GetHeader("Content-Encoding") == "gzip" {
		h.logger.Info("Content-Encoding: gzip")
		reader, err = gzip.NewReader(c.Request.Body)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "Invalid gzip body", err)
			return
		}
		defer reader.Close()
	} else {
		h.logger.Info("Content-Encoding: identity (not gzip)")
		reader = c.Request.Body
		defer reader.Close()
	}

	payload, err := io.ReadAll(reader)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid body", err)
		return
	}

	h.logger.Infof("Read payload of %d bits", len(payload))

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		h.logger.Errorf("Kafka write failed: %v", err)
		response.RespondError(c, http.StatusInternalServerError, "Failed to write message to Kafka", err)
		return
	}

	h.logger.Info("Kafka write successful")
	response.RespondSuccessWithData(c, nil)
}

func (h *LogHandler) UploadEvents(c *gin.Context) {
	writer, ok := h.kafkaWriters["events"]
	if !ok {
		response.RespondError(c, http.StatusInternalServerError, "Kafka writer not available for events", nil)
		return
	}

	var reader io.ReadCloser
	var err error

	if c.GetHeader("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(c.Request.Body)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "Invalid gzip body", err)
			return
		}
		defer reader.Close()
	} else {
		reader = c.Request.Body
		defer reader.Close()
	}

	payload, err := io.ReadAll(reader)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid body", err)
		return
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to write event to Kafka", err)
		return
	}

	response.RespondSuccessWithData(c, nil)
}

func (h *LogHandler) UploadContext(c *gin.Context) {
	writer, ok := h.kafkaWriters["context"]
	if !ok {
		response.RespondError(c, http.StatusInternalServerError, "Kafka writer not available for context", nil)
		return
	}

	var reader io.ReadCloser
	var err error

	if c.GetHeader("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(c.Request.Body)
		if err != nil {
			response.RespondError(c, http.StatusBadRequest, "Invalid gzip body", err)
			return
		}
		defer reader.Close()
	} else {
		reader = c.Request.Body
		defer reader.Close()
	}

	payload, err := io.ReadAll(reader)
	if err != nil {
		response.RespondError(c, http.StatusBadRequest, "Invalid body", err)
		return
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: payload,
	})
	if err != nil {
		response.RespondError(c, http.StatusInternalServerError, "Failed to write context to Kafka", err)
		return
	}

	response.RespondSuccessWithData(c, nil)
}
