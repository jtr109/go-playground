// [Search Insert Position - LeetCode](https://leetcode.com/problems/search-insert-position/submissions/)
// [代码随想录](https://programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html#%E7%9B%B8%E5%85%B3%E9%A2%98%E7%9B%AE%E6%8E%A8%E8%8D%90)

package searchinsertposition

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		// move cursors
		mid := left + (right-left)/2
		diff := target - nums[mid]
		if diff < 0 {
			// the target number should be the left part
			right = mid - 1
		} else if diff > 0 {
			// the target number should be the right part
			left = mid + 1
		} else {
			// find the target number
			return mid
		}
	}
	// left == right
	if target <= nums[left] {
		return left
	} else {
		return left + 1
	}
}
