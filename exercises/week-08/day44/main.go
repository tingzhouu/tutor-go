package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

type result struct {
	Error      error
	StatusCode int
	Url        string
}

func checkWebsites(urls []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)

	for _, url := range urls {
		g.Go(func() error {
			req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
			res, err := http.DefaultClient.Do(req)

			if err != nil {
				fmt.Printf("error: %s - %v\n", url, err)
				return err
			}
			fmt.Printf("%d - %s\n", res.StatusCode, url)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("some checks failed")
	}
}

func main() {
	checkWebsites([]string{
		"https://go.dev",
		"https://google.com",
		"https://thisurldoesnotexist12345.com",
	})
}
