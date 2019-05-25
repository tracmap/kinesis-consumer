package consumer

import "github.com/aws/aws-sdk-go/service/kinesis/kinesisiface"

// Option is used to override defaults when creating a new Consumer
type Option func(*Consumer)

// WithStorage overrides the default storage
func WithStorage(storage Storage) Option {
	return func(c *Consumer) {
		c.storage = storage
	}
}

// WithLogger overrides the default logger
func WithLogger(logger Logger) Option {
	return func(c *Consumer) {
		c.logger = logger
	}
}

// WithCounter overrides the default counter
func WithCounter(counter Counter) Option {
	return func(c *Consumer) {
		c.counter = counter
	}
}

// WithClient overrides the default client
func WithClient(client kinesisiface.KinesisAPI) Option {
	return func(c *Consumer) {
		c.client = client
	}
}

// ShardIteratorType overrides the starting point for the consumer
func WithShardIteratorType(t string) Option {
	return func(c *Consumer) {
		c.initialShardIteratorType = t
	}
}
