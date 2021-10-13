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
