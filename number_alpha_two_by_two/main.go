package main

import "fmt"

func main() {
	numCh := make(chan struct{})
	alphaCh := make(chan struct{})
	ch := make(chan struct{})

	go func() {
		for a := 'A'; a <= 'Z'; {
			<-alphaCh
			fmt.Print(string(a))
			a++
			fmt.Print(string(a))
			a++
			numCh <- struct{}{}
		}
		close(numCh)
		<-alphaCh // waste one element
	}()

	go func() {
		i := 1
		for range numCh {
			fmt.Print(i)
			i++
			fmt.Print(i)
			i++
			alphaCh <- struct{}{}
		}
		close(alphaCh)
		ch <- struct{}{}
	}()

	numCh <- struct{}{}
	<-ch
	close(ch)
}
