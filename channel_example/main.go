package main

func main() {
	c := make(chan int)
	c <- 1
	<-c
	c = nil
	c <- 2
}
