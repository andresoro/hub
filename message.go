package hub

import "time"

// Message is what is passed from publisher to subscriber
type Message struct {
	Topic     string `json:"topic"`
	Message   string `json:"message"`
	timeStamp time.Time
}

// NewMessage returns a new message
func NewMessage(topic, msg string) *Message {
	return &Message{
		Topic:     topic,
		Message:   msg,
		timeStamp: time.Now(),
	}
}

// Content returns a message
func (m *Message) Content() string {
	return m.Message
}
