package main

import (
	"fmt"
	"time"
)

func SendingBlockedWhenBufferIsFull() {
	ch := make(chan struct{}, 1)
	ch <- struct{}{}
	// buffer is full
	go func() {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("The data in channel will be spent.")
		<-ch
	}()
	ch <- struct{}{}
	fmt.Println("New data can be pushed now.")
}

func SendingToAClosedChannel() {
	ch := make(chan struct{})
	close(ch)
	ch <- struct{}{}
}

func SendingOnANilChannel() {
	ch := make(chan struct{})
	close(ch)
	ch = nil
	ch <- struct{}{}
}

func main() {
	// SendingBlockedWhenBufferIsFull()
	// SendingToAClosedChannel()
	SendingOnANilChannel()
}
