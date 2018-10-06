package main

import (
	"fmt"
	"time"

	"github.com/andresoro/hub"
)

func main() {
	hub := hub.New()

	hub.NewTopic("test")
	hub.NewTopic("test2")

	c, err := hub.Subscribe("test")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		for i := 0; i < 10; i++ {
			hub.Publish("test", []byte("test"))
			time.Sleep(1 * time.Second)
		}
		hub.Close("test")
	}()

	for msg := range c {
		fmt.Println(string(msg))
	}

}
