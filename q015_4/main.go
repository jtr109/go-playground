package main

import (
	"fmt"
	"sync"
	"time"
)

type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func()) {
	if o.done == 1 {
		return
	}
	fmt.Println("before lock")
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		o.done = 1
		f()
	}
}

func main() {
	o := &Once{}
	for i := 0; i < 1000; i++ {
		go o.Do(func() {
			fmt.Println("foo")
		})
	}
	time.Sleep(10 * time.Second)
}
