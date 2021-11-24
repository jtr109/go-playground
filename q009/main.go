// https://github.com/lifei6671/interview-go/blob/master/question/q009.md

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	ch := make(chan int)
	w := sync.WaitGroup{}
	w.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- rand.Int()
		}
		close(ch)
		w.Done()
	}()
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		w.Done()
	}()
	w.Wait()
}
