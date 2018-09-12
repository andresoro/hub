package hub

import "time"

// Message is what is passed from publisher to subscriber
type Message struct {
	topic     string
	message   string
	timeStamp time.Time
}

// NewMessage returns a new message
func NewMessage(topic, msg string) *Message {
	return &Message{
		topic:     topic,
		message:   msg,
		timeStamp: time.Now(),
	}
}

// Content returns a message
func (m *Message) Content() string {
	return m.message
}
