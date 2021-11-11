// https://github.com/lifei6671/interview-go/blob/master/question/q001.md

package main

import "fmt"

func printNumbers(numbers chan bool, letters chan bool) {
	i := 0
	for <-numbers {
		fmt.Print(i)
		i++
		fmt.Print(i)
		i++
		letters <- true
	}
}

func printLetters(numbers chan bool, letters chan bool, wait chan bool) {
	i := 'a'
	for <-letters {
		if i > 'z' {
			close(wait)
			return
		}
		fmt.Print(string(i))
		i++
		fmt.Print(string(i))
		i++
		numbers <- true
	}
}

func main() {
	numbers := make(chan bool)
	letters := make(chan bool)
	wait := make(chan bool)
	go printNumbers(numbers, letters)
	go printLetters(numbers, letters, wait)
	numbers <- true
	<-wait
}
