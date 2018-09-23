package hub

import (
	"sync"
)

// Subs slice of sub
type Subs []*Sub

// Sub represents a subscriber on a given Hub
type Sub struct {
	topics  []string
	channel chan *Message
	lock    sync.Mutex
}

// NewSub returns a new sub
func NewSub(topics ...string) *Sub {
	var t []string
	for _, topic := range topics {
		t = append(t, topic)
	}

	return &Sub{
		topics:  t,
		channel: make(chan *Message),
	}
}

// Subscribe to another topic
func (s *Sub) Subscribe(topics ...string) {
	for _, topic := range topics {
		s.topics = append(s.topics, topic)
	}
}

// Channel returns a subscribers read channel
func (s *Sub) Channel() chan *Message {
	return s.channel
}

// Close a subs message channel
func (s *Sub) Close() {
	s.lock.Lock()
	close(s.channel)
	s.lock.Unlock()
}
