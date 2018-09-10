package hub

import "errors"

// Hub is a message broker (pub-sub) that relays messages between publishers and subscribers
// publishers write to a topic
type Hub struct {
	topics     map[string]chan string
	bufferSize int
}

// NewHub returns a Hub with a channel buffer size
func NewHub(bufferSize int) *Hub {
	return &Hub{
		topics:     make(map[string]chan string),
		bufferSize: bufferSize,
	}
}

// NewTopic adds a topic and buffered channel to the Hub. Returns an error if topic exists
func (h *Hub) NewTopic(name string) error {

	if h.exist(name) {
		return errors.New("Topic already exists")
	}

	h.topics[name] = make(chan string, h.bufferSize)
	return nil

}

//Publish content to a topic
func (h *Hub) Publish(content, topic string) error {

	if !h.exist(topic) {
		return errors.New("Topic does not exist")
	}

	h.topics[topic] <- content
	return nil

}

// Fetch will return last string that was published
func (h *Hub) Fetch(topic string) (string, error) {

	temp := make(chan string)
	go func(c chan string) {
		for msg := range h.topics[topic] {
			c <- msg
		}
		close(c)
	}(temp)
	return <-temp, nil
}

func (h *Hub) exist(topic string) bool {
	if _, ok := h.topics[topic]; ok {
		return true
	}
	return false
}
