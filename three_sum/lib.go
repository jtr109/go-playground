// https://leetcode.com/problems/3sum/

package threesum

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			// move until nums[i] != nums[i-1]
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
			}
			if sum < 0 {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			} else {
				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}
	return result
}

// time complexity: O(n)
