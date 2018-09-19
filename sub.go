package hub

// Sub represents a subscriber on a given Hub
type Sub struct {
	topics  []string
	channel chan *Message
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
