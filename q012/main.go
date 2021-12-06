package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			func() { // learn: recover at the end of current function
				defer func() {
					if err := recover(); err != nil {
						fmt.Println(err)
					}
				}()
				time.Sleep(time.Second)
				proc()
			}()
		}
	}()

	select {}
}

func proc() {
	panic("ok")
}
