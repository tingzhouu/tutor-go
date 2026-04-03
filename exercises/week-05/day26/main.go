package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func downloadFile(ctx context.Context, name string, delay time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Printf("%s cancelled\n", name)
		return
	case <-time.After(delay):
		fmt.Printf("%s finished\n", name)
	}
}

func main() {
	downloadItems := []struct {
		filename string
		delay    time.Duration
	}{{"file1.txt", time.Millisecond * 300},
		{"file2.txt", time.Millisecond * 100},
		{"file3.txt", time.Millisecond * 700},
		{"file4.txt", time.Millisecond * 200},
		{"file5.txt", time.Millisecond * 4000},
	}

	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	for _, item := range downloadItems {
		wg.Add(1)
		go func() {
			defer wg.Done()
			downloadFile(ctx, item.filename, item.delay)
		}()
	}

	wg.Wait()
	fmt.Println("all downloads complete")
}
