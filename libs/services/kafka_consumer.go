package baseServices

import (
	"context"
	"fmt"
	baseConstants "github.com/quangtran88/anifni-base/libs/constants"
	basePorts "github.com/quangtran88/anifni-base/libs/ports"
	baseUtils "github.com/quangtran88/anifni-base/libs/utils"
	"github.com/segmentio/kafka-go"
	"log"
	"strings"
)

type KafkaConsumer struct {
	hosts []string
}

func NewKafkaConsumer() *KafkaConsumer {
	env := baseUtils.GetEnvManager()
	kafkaHostsEnv := env.GetEnv(baseConstants.KafkaHostEnvKey)
	hosts := strings.Split(kafkaHostsEnv, ",")
	log.Printf("Init Kafka Consumer with hosts %s", hosts)
	return &KafkaConsumer{hosts}
}

func (c KafkaConsumer) ConsumeAsync(topic string, group string, consumer basePorts.ConsumerHandler) {
	r := c.initReader(topic, group)
	go c.consumeMessage(context.Background(), r, consumer)
}

func (c KafkaConsumer) Consume(topic string, group string, consumer basePorts.ConsumerHandler) {
	r := c.initReader(topic, group)
	c.consumeMessage(context.Background(), r, consumer)
}

func (c KafkaConsumer) initReader(topic string, group string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  c.hosts,
		GroupID:  group,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
	})
}

func (c KafkaConsumer) consumeMessage(ctx context.Context, r *kafka.Reader, consumer basePorts.ConsumerHandler) {
	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Print(err)
			break
		}
		log.Printf("Consume kafka message %s", c.serializeMessage(msg))
		consumer(basePorts.EventMessage{
			Key:   string(msg.Key),
			Value: string(msg.Value),
		})
	}

	if err := r.Close(); err != nil {
		log.Printf("Failed to close reader: %v", err)
	}
}

func (c KafkaConsumer) serializeMessage(msg kafka.Message) string {
	s := make([]string, 0)
	s = append(s, fmt.Sprintf("T/P/O: %s/%v/%v", msg.Topic, msg.Partition, msg.Offset))
	s = append(s, fmt.Sprintf("Key: %s", msg.Key))
	s = append(s, fmt.Sprintf("Value: %s", msg.Value))
	s = append(s, fmt.Sprintf("Header: %s", c.serializeHeaders(msg.Headers...)))
	return strings.Join(s, " - ")
}

func (c KafkaConsumer) serializeHeaders(messages ...kafka.Header) string {
	s := make([]string, 0, len(messages))
	for _, msg := range messages {
		s = append(s, fmt.Sprintf("\"%s\" : \"%s\"", msg.Key, msg.Value))
	}
	return "{ " + strings.Join(s, ", ") + " }"
}
