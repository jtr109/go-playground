// https://github.com/lifei6671/interview-go/blob/master/question/q009.md

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func generate(wg *sync.WaitGroup, ch chan int, count int) {
	defer wg.Done()
	for i := 0; i < count; i++ {
		ch <- rand.Int()
	}
	close(ch)
}

func print(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go generate(&wg, ch, 5)
	go print(&wg, ch)
	wg.Wait()
}
