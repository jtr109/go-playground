package main

import (
	"fmt"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		c <- 10
		close(c)
	}()
	for {
		select {
		case x, ok := <-c:
			fmt.Printf("Retrieved at %v: x=%v, ok=%v\n", time.Now().Format(timeFormat), x, ok)
			time.Sleep(500 * time.Millisecond)
		default:
			fmt.Printf("Not retrieved at %v\n", time.Now().Format(timeFormat))
			time.Sleep(500 * time.Millisecond)
		}
	}
}
