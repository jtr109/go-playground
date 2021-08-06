// https://www.codewars.com/kata/555615a77ebc7c2c8a0000b8/train/go

package vasyaclerk

import "fmt"

func Tickets(peopleInLine []int) string {
	balance := make(map[int]int)
	fmt.Println(peopleInLine)
	for _, d := range peopleInLine {
		switch d {
		case 25:
		case 50:
			if balance[25] == 0 {
				return "NO"
			}
			balance[25]--
		case 100:
			if balance[50] > 0 && balance[25] > 0 {
				balance[25]--
				balance[50]--
				balance[100]++
			} else if balance[25] > 2 {
				balance[25] -= 3
			} else {
				return "NO"
			}
		}
		balance[d]++
	}
	return "YES"
}
