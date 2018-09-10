package hub

import (
	"fmt"
	"testing"
)

func TestNewTopic(t *testing.T) {
	hub := NewHub(5)
	var err error

	_ = hub.NewTopic("test")

	err = hub.NewTopic("test2")
	if err != nil {
		t.Error("Should not return error on unique topic")
	}

	err = hub.NewTopic("test")
	if err == nil {
		t.Error(err)
	}

}

func TestPublish(t *testing.T) {
	hub := NewHub(5)

	hub.NewTopic("topic")
	in := "input"
	hub.Publish(in, "topic")

	out, _ := hub.Fetch("topic")

	if out != in {
		fmt.Print(out)
		t.Error("Subscribe is not returning valued published")
	}

}
