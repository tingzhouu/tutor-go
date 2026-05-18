package main

import (
	"fmt"
	"time"
	"tutor-go/projects/pubsub/broker"
)

func main() {
	b := broker.NewBroker()
	sportsCh := b.Subscribe("sports")

	go func() {
		for msg := range sportsCh {
			time.Sleep(time.Millisecond * 500)
			fmt.Println(time.Now().Format("15:04:05.000"), "got:", msg)
		}
	}()

	for i := range 30 {
		fmt.Println(time.Now().Format("15:04:05.000"), "publish", i)
		b.Publish("sports", fmt.Sprintf("msg %d", i))
	}

	time.Sleep(time.Millisecond * 10000)
}
