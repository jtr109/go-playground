package main

import (
	"context"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func find(nums []int, target int, timeout time.Duration) bool {
	fmt.Printf("finding %d\n", target)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan int)
	defer close(ch)
	groupLength := int(math.Ceil(float64(len(nums)) / 10.0))
	for i := 0; i < len(nums); i = i + groupLength {
		go func(j int) {
			for k := 0; k < groupLength; k++ {
				fmt.Printf("Comparing number in index %d in routine %d.\n", j+k, j)
				select {
				case <-ctx.Done():
					fmt.Println("canceled.")
					return
				default:
					if nums[j+k] == target {
						fmt.Printf("found %d in routine %d!\n", nums[j+k], j)
						cancel()
						ch <- j + k
					}
				}
			}
		}(i)
	}
	select {
	case <-time.After(timeout):
		return false
	case <-ch:
		return true
	}
}

func main() {
	count := 100000
	s := []int{}
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < count; i++ {
		s = append(s, r1.Intn(count))
	}
	result := find(s, r1.Intn(count), 3*time.Second)
	fmt.Println(result)
	time.Sleep(10 * time.Millisecond)
}
