package basePorts

import "context"

type EventProducer interface {
	Produce(ctx context.Context, topic string, key string, value string) error
	ProduceMultiple(ctx context.Context, topic string, messages []EventMessage) error
}
