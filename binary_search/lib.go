// https://programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html#_704-%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE

package binarysearch

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		diff := nums[mid] - target
		if diff < 0 {
			// the number indexed in mid value is less than target
			left = mid + 1
		} else if diff > 0 {
			// the number indexed in mid value is greater than target
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 以下代码和上面有两个差异：使用了递归、后一个下标不包含在比较列表中
// 使用包含的下标能带来的好处是减少了比较的成本，因为最终条件会收束到一个明确的下标，这个下标对应的值如果符合目标值，则找到了目标
// 如果不符合目标，则目标值不存在
// 而使用递归和使用迭代在这个例子中的复杂度是一样的，都是 O(logN)

// 后记：在这种左闭右开的方法中，因为 num[mid] 已经和 target 比较过了，所以如果应该选择右区间时，也不应该再引入 mid 下标了，所以 start = mid + 1 更合适

// func search(nums []int, target int) int {
// 	return binarySearch(nums, target, 0, len(nums))
// }

// func binarySearch(nums []int, target int, start int, end int) int {
// 	if start == end {
// 		// there is no number to be compared
// 		return -1
// 	}
// 	if start+1 == end {
// 		// the number indexed in start value is the only one
// 		if nums[start] == target {
// 			return start
// 		} else {
// 			return -1
// 		}
// 	}
// 	mid := (start + end) / 2
// 	diff := nums[mid] - target
// 	if diff < 0 {
// 		// the mid number is less than the target
// 		return binarySearch(nums, target, mid, end)
// 	} else if diff > 0 {
// 		// the mid number is greater than the target
// 		return binarySearch(nums, target, start, mid)
// 	} else {
// 		// the mid number matches the target
// 		return mid
// 	}
// }
