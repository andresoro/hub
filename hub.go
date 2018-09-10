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

	err := h.exist(name)
	if err != nil {
		return err
	}

	h.topics[name] = make(chan string, h.bufferSize)
	return nil

}

//Publish content to a topic
func (h *Hub) Publish(content, topic string) error {

	err := h.exist(topic)
	if err != nil {
		return err
	}

	h.topics[topic] <- content
	return nil

}

// Subscribe will return last string that was published
func (h *Hub) Subscribe(topic string) (string, error) {
	err := h.exist(topic)
	if err != nil {
		return "", err
	}

	return <-h.topics[topic], nil

}

func (h *Hub) exist(topic string) error {
	if _, ok := h.topics[topic]; ok {
		return errors.New("Topic already exists")
	}
	return nil
}
