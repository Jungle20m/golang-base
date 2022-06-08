package utils

import (
	"strings"

	"github.com/segmentio/kafka-go"
)

type KafkaInstance struct {
	Brokers string
}

func NewKafkaInstance(kafkaBroker string) *KafkaInstance {
	return &KafkaInstance{
		Brokers: kafkaBroker,
	}
}

func (k *KafkaInstance) NewWriter() (*kafka.Writer, error) {
	return &kafka.Writer{
		Addr:     kafka.TCP(k.Brokers),
		Balancer: &kafka.LeastBytes{},
	}, nil
}

func (k *KafkaInstance) NewReader(topic string, params ...string) (*kafka.Reader, error) {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  strings.Split(k.Brokers, ","),
		Topic:    topic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	}), nil
}
