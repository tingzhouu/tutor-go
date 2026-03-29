package main

import (
	"fmt"
	"time"
)

func fetchPrice(item string) int {
	time.Sleep(time.Second * 1)
	prices := map[string]int{
		"shirt": 1052,
		"pants": 5712,
		"shoes": 9999,
	}
	return prices[item]
}

type result struct {
	price int
	item  string
}

func main() {
	ch := make(chan result)
	items := []string{"shirt", "pants", "shoes"}

	for _, item := range items {
		go func() {
			ch <- result{fetchPrice(item), item}
		}()
	}

	for range items {
		r := <-ch
		fmt.Printf("%s %d\n", r.item, r.price)
	}
}
