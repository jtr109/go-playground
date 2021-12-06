package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	sync.Mutex
}

func NewBan() *Ban {
	o := &Ban{visitIPs: make(map[string]time.Time)}
	return o
}

func (o *Ban) visit(ip string) bool {
	o.Lock()
	defer o.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	success := int64(0)

	ban := NewBan()

	wait := &sync.WaitGroup{}

	wait.Add(1000 * 100)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wait.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j) // learn: send j as parameter is important
		}

	}
	wait.Wait()

	fmt.Println("success:", success)
}
