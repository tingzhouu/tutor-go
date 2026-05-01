package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func slowFunc(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Printf("cancelled\n")
		return
	case <-time.After(time.Second * 3):
		fmt.Printf("slow func completed\n")
	}
	fmt.Println("doing cleanup work")
}

func runSlowFuncAndCancelOnTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup

	for range 3 {
		wg.Go(func() {
			slowFunc(ctx)
		})
	}

	time.Sleep(1 * time.Second)

	wg.Wait()
	fmt.Printf("cancelled reason: %v\n", ctx.Err())
}

func fetchUrls(ctx context.Context) {
	var wg sync.WaitGroup
	urls := []struct {
		url     string
		seconds int
	}{
		{url: "https://google.com", seconds: 1},
		{url: "https://alibaba.com", seconds: 2},
		{url: "http://localhost:3000/delay", seconds: 3},
	}

	for _, item := range urls {
		wg.Go(func() {
			req, err := http.NewRequestWithContext(ctx, "GET", item.url, nil)
			if err != nil {
				fmt.Printf("request error: %v\n", err)
				return
			}

			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				fmt.Printf("error occured - %v\n", err)
				return
			}
			fmt.Printf("%s - successful for url %s\n", resp.Status, item.url)
		})
	}

	wg.Wait()
}

func main() {
	fmt.Println("main-start")
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	var wg sync.WaitGroup
	defer cancel()

	wg.Go(func() {
		fetchUrls(ctx)
	})

	wg.Wait()
}
