package main

import (
	"fmt"
	"time"
)

func main() {

	// ch := make(chan struct{}) // unbuffered channel
	ch := make(chan struct{}, 3) // buffered channel

	for i := range 3 {
		go func() {
			time.Sleep(time.Second * time.Duration(i+1))
			fmt.Printf("Worker %d done\n", i)
			fmt.Printf("Worker %d about to send\n", i)
			ch <- struct{}{}
			fmt.Printf("Worker %d sent\n", i)
		}()
	}

	time.Sleep(time.Second * 5)

	for range 3 {
		<-ch
	}

	time.Sleep(time.Millisecond * 100)
}
