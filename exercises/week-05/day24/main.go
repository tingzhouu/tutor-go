package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchItem(item string) int {
	if item == "pants" {
		time.Sleep(time.Second * 3)
	}
	priceList := map[string]int{
		"pants": 100,
		"shirt": 300,
		"shoes": 500,
	}

	return priceList[item]
}

type result struct {
	item  string
	price int
}

func main() {
	ch := make(chan result)
	var wg sync.WaitGroup
	items := []string{"shirt", "pants", "shoes"}
	for _, item := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := fetchItem(item)
			ch <- result{item, res}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		fmt.Printf("%s %d\n", r.item, r.price)
	}
}
