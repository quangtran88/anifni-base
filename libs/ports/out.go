package basePorts

import (
	"context"
)

type EventProducer interface {
	Produce(ctx context.Context, topic string, key string, value string) error
	ProduceMultiple(ctx context.Context, topic string, messages []EventMessage) error
}

type ConsumerHandler func(message EventMessage)

type EventConsumer interface {
	ConsumeAsync(topic string, group string, consumer ConsumerHandler)
	Consume(topic string, group string, consumer ConsumerHandler)
}
