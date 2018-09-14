package hub

// Sub represents a subscriber on a given Hub
type Sub struct {
	topics []string
}

// NewSub returns a new sub
func NewSub(topics ...string) *Sub {
	var t []string
	for _, topic := range topics {
		t = append(t, topic)
	}

	return &Sub{
		topics: t,
	}
}
