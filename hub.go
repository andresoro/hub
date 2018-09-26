package hub

import (
	"errors"
	"sync"
	"time"
)

// Hub is a publish/subscribe broker
type Hub struct {
	lock   sync.Mutex
	topics map[string]chan []byte
}

type message struct {
	topic   string
	content []byte
	time    time.Time
}

// New returns a new Hub and starts go routine to handle incoming messages
func New() *Hub {
	h := &Hub{
		topics: make(map[string]chan []byte),
	}

	return h
}

// Publish a message to topic
func (h *Hub) Publish(topic string, content []byte) error {

	msg := message{
		topic:   topic,
		content: content,
		time:    time.Now(),
	}

	if c, ok := h.topics[topic]; ok {
		c <- msg.content
		return nil
	}
	return errors.New("Topic does not exist")
}

// PublishNew will publish to a topic channel or create a new one if it doesnt exist
func (h *Hub) PublishNew(topic string, content []byte) {
	h.NewTopic(topic)
	h.Publish(topic, content)
}

// NewTopic adds a topic to the hub
func (h *Hub) NewTopic(topic string) error {
	h.lock.Lock()
	defer h.lock.Unlock()
	if _, ok := h.topics[topic]; ok {
		return errors.New("Topic already exists")
	}
	h.topics[topic] = make(chan []byte)
	return nil
}

func (h *Hub) Close(topic string) {
	close(h.topics[topic])
}

// Subscribe returns a read only channel to receive messages
func (h *Hub) Subscribe(topic string) (<-chan []byte, error) {
	if _, ok := h.topics[topic]; !ok {
		return nil, errors.New("Topic does not exist")
	}

	return h.topics[topic], nil
}
