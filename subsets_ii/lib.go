package subsetsii

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	// cache to store subsets, using the length of subsets as key
	sort.Ints(nums)
	result, current := [][]int{}, [][]int{}
	result = append(result, []int{})
	for i, n := range nums {
		fmt.Println("n", n)
		newCurrent := [][]int{}
		if i != 0 && nums[i-1] == n {
			// when n is equal to the last one,
			// the new result of appending to last result is duplicated,
			// but appending to last current is new
			for _, item := range current {
				newCurrent = append(newCurrent, append(item, n))
			}
		} else {
			for _, item := range result {
				fmt.Println("result in loop", result)
				fmt.Println("item", item)
				newCurrent = append(newCurrent, append(item, n))
				fmt.Println("result after append", result)
				fmt.Println("---")
			}
		}
		current = newCurrent
		result = append(result, current...)
	}
	return result
}
