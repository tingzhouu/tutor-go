package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func checkApi() string {
	time.Sleep(time.Millisecond * 200)
	return "healthy"
}

func checkDatabase() string {
	time.Sleep(time.Millisecond * 800)
	return "healthy"
}

func checkCache() string {
	time.Sleep(time.Millisecond * 100)
	return "error: connection refused"
}

func checkQueue() string {
	time.Sleep(time.Millisecond * 5000)
	return "healthy"
}

type itemCheck struct {
	system  string
	checkFn func() string
}

type res struct {
	system string
	result string
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan res)
	checklist := []itemCheck{{"api", checkApi}, {"database", checkDatabase}, {"cache", checkCache}, {"queue", checkQueue}}

	for _, item := range checklist {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- res{item.system, item.checkFn()}
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		fmt.Printf("%s-%s\n", r.system, r.result)
	}

	fmt.Println("goroutines still running:", runtime.NumGoroutine())
}
