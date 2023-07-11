package consumer

import (
	"github.com/author_name/project_name/pkg/pubsub/publisher"
)

type Publisher interface {
	RegisterSubscription(topic string) publisher.MessageChannel
}

type Consumer struct {
	publisher Publisher
}

func New(publisher Publisher) *Consumer {
	return &Consumer{
		publisher: publisher,
	}
}

func (c *Consumer) SubscribeTopic(topic string, handler func(message publisher.Message)) {
	subscription := c.publisher.RegisterSubscription(topic)
	go func(channel publisher.MessageChannel) {
		for {
			message := <-channel
			handler(message)
		}
	}(subscription)
}

func (c *Consumer) Unsubscribe() {
	// TODO implement unsubscribe
}
