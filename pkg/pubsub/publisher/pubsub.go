package publisher

import (
	"sync"
)

type Publisher struct {
	locker           *sync.RWMutex
	messageQueue     MessageChannel
	subscribeChannel map[string][]MessageChannel
}

func New() *Publisher {
	return &Publisher{
		locker:           new(sync.RWMutex),
		messageQueue:     make(MessageChannel, 10000),
		subscribeChannel: make(map[string][]MessageChannel),
	}
}

func (pb *Publisher) Run() {
	go func() {
		for {
			message := <-pb.messageQueue

			subscribeChannel, ok := pb.subscribeChannel[message.Topic]
			if ok {
				for _, channel := range subscribeChannel {
					channel <- message
				}
			}
		}
	}()
}

func (pb *Publisher) Publish(topic string, content interface{}) {
	pb.messageQueue <- *NewMessage(topic, content)
}

func (pb *Publisher) RegisterSubscription(topic string) MessageChannel {
	newChannelForSubscriber := make(MessageChannel)
	pb.locker.Lock()
	currentSubscriberChannel, ok := pb.subscribeChannel[topic]
	if ok {
		currentSubscriberChannel = append(currentSubscriberChannel, newChannelForSubscriber)
		pb.subscribeChannel[topic] = currentSubscriberChannel
	} else {
		pb.subscribeChannel[topic] = []MessageChannel{newChannelForSubscriber}
	}
	pb.locker.Unlock()
	return newChannelForSubscriber
}
