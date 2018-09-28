package hub

import (
	"testing"
)

func TestHub(t *testing.T) {

	hub := New()

	//test topic uniqueness
	hub.NewTopic("test")
	err := hub.NewTopic("test")
	if err == nil {
		t.Error("NewTopic should return error is topic exists")
	}

	//test subscribing to topic, ignore return chan
	_, err = hub.Subscribe("test")
	if err != nil {
		t.Error("Subscribe should to return error on valid topic")
	}

	_, err = hub.Subscribe("topic")
	if err == nil {
		t.Error("Subscribe should return error on non existant topic")
	}

}
