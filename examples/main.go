package main

import (
	"fmt"
	"time"

	"github.com/andresoro/hub"
)

func main() {
	hub := hub.New()

	hub.NewTopic("test")
	c, err := hub.Subscribe("test")
	if err != nil {
		fmt.Println(err)
	}

	go func() {
		for i := 0; i < 10; i++ {
			hub.Publish("test", []byte("test"))
			time.Sleep(2 * time.Second)
		}
		hub.Close("test")
	}()

	for msg := range c {
		fmt.Println(string(msg))
	}

}
