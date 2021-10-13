package binarysearch

func search(nums []int, target int) int {
	return binarySearch(nums, target, 0, len(nums))
}

func binarySearch(nums []int, target int, start int, end int) int {
	if start == end {
		// there is no number to be compared
		return -1
	}
	if start+1 == end {
		// the number indexed in start value is the only one
		if nums[start] == target {
			return start
		} else {
			return -1
		}
	}
	mid := (start + end) / 2
	diff := nums[mid] - target
	if diff < 0 {
		// the mid number is less than the target
		return binarySearch(nums, target, mid, end)
	} else if diff > 0 {
		// the mid number is greater than the target
		return binarySearch(nums, target, start, mid)
	} else {
		// the mid number matches the target
		return mid
	}
}
