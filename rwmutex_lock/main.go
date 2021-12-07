package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := new(sync.RWMutex)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func(j int) {
			dur := 100 * j
			time.Sleep(time.Duration(dur) * time.Millisecond)
			fmt.Printf("Reader %d had slept %d milliseconds.\n", j, dur)
			mu.RLock()
			fmt.Printf("Reader %d got a read lock.\n", j)
			time.Sleep(time.Duration(100*j) * time.Millisecond)
			mu.RUnlock()
			fmt.Printf("Reader %d released a read lock\n", j)
			wg.Done()
		}(i)
		wg.Add(1)
	}
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Writer had slept 500 milliseconds.")
		mu.Lock()
		fmt.Println("Writer got the lock.")
		time.Sleep(200 * time.Millisecond)
		mu.Unlock()
		fmt.Println("Writer released the lock.")
		wg.Done()
	}()
	wg.Add(1)
	wg.Wait()
}
