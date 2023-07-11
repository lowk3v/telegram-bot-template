package publisher

type MessageChannel chan Message

type Message struct {
	Topic   string
	Content interface{}
}

func NewMessage(topic string, content interface{}) *Message {
	return &Message{
		Topic:   topic,
		Content: content,
	}
}

func (m *Message) GetTopic() string {
	return m.Topic
}

func (m *Message) GetContent() interface{} {
	return m.Content
}
