package kafka_mem

import (
	"time"

	"github.com/adarsh-devappsys/rest-with-protobuf-logs/internal/app/config"
	"github.com/segmentio/kafka-go"
)

func InitKafkaWriters(cfg *config.TKafkaConfig) (map[string]*kafka.Writer, error) {
	writers := make(map[string]*kafka.Writer)

	// Map Acks string to Kafka RequiredAcks enum
	ack, ok := map[string]kafka.RequiredAcks{
		"none": kafka.RequireNone,
		"one":  kafka.RequireOne,
		"all":  kafka.RequireAll,
	}[cfg.Acks]
	if !ok {
		ack = kafka.RequireAll
	}

	// Create Kafka writers for each topic
	for name, topic := range cfg.Topics {
		writer := &kafka.Writer{
			Addr:         kafka.TCP(cfg.Brokers...), // Use the brokers from the config
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: ack,
			Async:        cfg.Async,
			WriteTimeout: time.Second * time.Duration(cfg.WriteTimeout),
			ReadTimeout:  time.Second * time.Duration(cfg.ReadTimeout),
		}
		writers[name] = writer
	}

	return writers, nil
}
