package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func increment() {
	atomic.AddInt64(&counter, 1)
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
