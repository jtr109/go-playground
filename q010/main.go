package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	rmx     *sync.RWMutex
	content map[string]*entity
}

type entity struct {
	ch  chan struct{} // use a channel to block receiver
	val interface{}
}

func (s *SafeMap) Out(key string, val interface{}) {
	s.rmx.Lock()
	defer s.rmx.Unlock()
	ent, ok := s.content[key]
	if !ok {
		ent = &entity{
			ch: make(chan struct{}),
		}
	}
	ent.val = val
	s.content[key] = ent
	close(ent.ch)
}

func (s *SafeMap) Rd(key string, timeout time.Duration) interface{} {
	s.rmx.RLock()
	ent, ok := s.content[key]
	if !ok {
		ent = &entity{
			ch: make(chan struct{}),
		}
		// learn: unlock the read-lock and lock a write-lock
		s.rmx.RUnlock()
		s.rmx.Lock()
		s.content[key] = ent
		s.rmx.Unlock()
	} else {
		s.rmx.RUnlock() // learn: don't forget to unlock if value exists
	}
	select {
	case <-ent.ch: // block until the channel is closed
		return ent.val
	case <-time.After(timeout):
		return nil
	}
}

func main() {
	safeMap := SafeMap{
		rmx:     new(sync.RWMutex),
		content: make(map[string]*entity),
	}
	key := "foo"
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(safeMap.Rd(key, 5*time.Second))
		}()
	}
	fmt.Println("sleep 1 second")
	time.Sleep(time.Second)
	fmt.Println("update map...")
	safeMap.Out(key, 1)
	time.Sleep(time.Second)
}
