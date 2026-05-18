package broker

import "fmt"

type Broker struct {
	topics map[string][]chan string
}

func NewBroker() *Broker {
	return &Broker{topics: map[string][]chan string{}}
}

func (b *Broker) Subscribe(topic string) <-chan string {
	ch := make(chan string, 16)
	b.topics[topic] = append(b.topics[topic], ch)
	return ch
}

func (b *Broker) Publish(topic, message string) {
	channels := b.topics[topic]
	for _, ch := range channels {
		// ch <- message

		select {
		case ch <- message:
		default:
			fmt.Println("dropped:", message)
		}
	}
}
