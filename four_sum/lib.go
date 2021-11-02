// https://leetcode.com/problems/4sum/submissions/

package foursum

import "sort"

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := len(nums) - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
				}
				if sum < target {
					// move left
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
				} else {
					// move right
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}
	return result
}
