package main

import (
	"fmt"
	"slices"
	"time"
)

type result struct {
	item  string
	price int
	err   error
}

func fetchPriceWithError(item string) (int, error) {
	if item == "shoes" {
		time.Sleep(time.Second * 3)
	} else if item == "pants" {
		return 0, fmt.Errorf("service unavailable")
	} else {
		time.Sleep(time.Millisecond * 500)

	}

	var priceList map[string]int = map[string]int{"shoes": 100, "pants": 500, "shirt": 2800}
	return priceList[item], nil
}

func main() {
	items := []string{"shoes", "pants", "shirt"}
	ch := make(chan result)

	for _, item := range items {
		go func() {
			price, err := fetchPriceWithError(item)
			ch <- result{item: item, price: price, err: err}
		}()
	}
	succeeded := []string{}
	failed := []string{}

	// received := 0
	timeout := time.After(time.Second * 1)
	// for received < len(items) {
	for range len(items) {
		select {
		case <-timeout:
			fmt.Printf("Did not receive all items\n")
			// received++

		case r := <-ch:
			if r.err == nil {
				succeeded = append(succeeded, r.item)
				fmt.Printf("%s, %d\n", r.item, r.price)
			} else {
				failed = append(failed, r.item)
				fmt.Printf("error occured for %s: %v\n", r.item, r.err)
			}

			// received++
		}
	}

	for _, item := range items {
		if slices.Contains(succeeded, item) {
			fmt.Printf("%s: success\n", item)
		} else if slices.Contains(failed, item) {
			fmt.Printf("%s: failure\n", item)
		} else {
			fmt.Printf("%s: timeout\n", item)

		}
	}
}
