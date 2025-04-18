package infra

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
}

func NewKafkaProducer(brokers []string, topic string) *KafkaProducer {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: brokers,
		Topic:   topic,
		Async:   true,
	})
	return &KafkaProducer{Writer: writer}
}

func (p *KafkaProducer) Publish(ctx context.Context, key, value []byte) error {
	err := p.Writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		log.Printf("Failed to publish message to Kafka: %v", err)
		return err
	}
	log.Printf("Message published to Kafka: %s", value)
	return nil
}

func (p *KafkaProducer) Close() error {
	return p.Writer.Close()
}
