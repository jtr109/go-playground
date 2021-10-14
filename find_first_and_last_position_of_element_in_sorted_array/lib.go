package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	return []int{
		searchLeft(nums, target),
		searchRight(nums, target),
	}
}

func searchLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			// two cursors overlap and matches the target number
			if nums[left] == target {
				return left
			} else {
				return -1
			}
		}

		mid := left + (right-left)/2
		diff := target - nums[mid]
		if diff < 0 {
			// the target is less than the mid value
			right = mid - 1
		} else if diff > 0 {
			// the target is greater than the mid value
			left = mid + 1
		} else {
			// the mid value matches, but may not be the leftmost number
			right = mid
		}
	}
	return -1
}
func searchRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		if left == right {
			if nums[left] == target {
				return left
			} else {
				return -1
			}
		}

		mid := left + (right-left)/2 + 1
		diff := target - nums[mid]
		if diff < 0 {
			// the target is less than the mid value
			right = mid - 1
		} else if diff > 0 {
			// the target is greater than the mid value
			left = mid + 1
		} else {
			// the mid value matches, but may not be the rightmost number
			left = mid
		}
	}
	return -1
}
