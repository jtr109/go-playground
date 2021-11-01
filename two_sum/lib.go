// https://leetcode.com/problems/two-sum/

package twosum

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)
	for i, n := range nums {
		if j, ok := cache[target-n]; ok {
			return []int{j, i}
		}
		cache[n] = i
	}
	return []int{}
}
