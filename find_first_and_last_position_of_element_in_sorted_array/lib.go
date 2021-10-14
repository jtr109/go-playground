// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

package findfirstandlastpositionofelementinsortedarray

func searchRange(nums []int, target int) []int {
	return []int{
		searchLeft(nums, target),
		searchRight(nums, target),
	}
}

// func searchLeft(nums []int, target int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		if left == right {
// 			// two cursors overlap and matches the target number
// 			if nums[left] == target {
// 				return left
// 			} else {
// 				return -1
// 			}
// 		}

// 		mid := left + (right-left)/2
// 		diff := target - nums[mid]
// 		if diff < 0 {
// 			// the target is less than the mid value
// 			right = mid - 1
// 		} else if diff > 0 {
// 			// the target is greater than the mid value
// 			left = mid + 1
// 		} else {
// 			// the mid value matches, but may not be the leftmost number
// 			right = mid
// 		}
// 	}
// 	return -1
// }
// func searchRight(nums []int, target int) int {
// 	left, right := 0, len(nums)-1
// 	for left <= right {
// 		if left == right {
// 			if nums[left] == target {
// 				return left
// 			} else {
// 				return -1
// 			}
// 		}

// 		mid := left + (right-left)/2 + 1
// 		diff := target - nums[mid]
// 		if diff < 0 {
// 			// the target is less than the mid value
// 			right = mid - 1
// 		} else if diff > 0 {
// 			// the target is greater than the mid value
// 			left = mid + 1
// 		} else {
// 			// the mid value matches, but may not be the rightmost number
// 			left = mid
// 		}
// 	}
// 	return -1
// }

// 上述方法的主要思路就是分别得到最左边和最右边的两个匹配元素的索引
// 但是存在一个问题是边界判断太混乱了
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/discuss/14734/Easy-java-O(logn)-solution
// 上面这个解法提供了一个思路，即定义一个预备结果 -1。
// 在查找最左元素时：
// 如果目标小于二分法中间数，那么右侧边界左移
// 如果目标大于二分法中间数，那么左侧边界右移
// 如果目标等于二分法中间数，那么中间数可能是最左侧的目标元素，也可能在目标元素的右侧，所以更新预备结果，并将右侧边界左移。
// 查找最右侧元素时
// 如果目标不等于中间数，操作和上述一致
// 如果目标等于二分法中间数，那么中间数可能是最右侧的目标元素，也可能在目标元素的左侧，所以更新预备结果，并将左侧边界右移。

// func searchRange(nums []int, target int) []int {}

func searchLeft(nums []int, target int) int {
	result := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		diff := target - nums[mid]
		if diff < 0 {
			right = mid - 1
		} else if diff > 0 {
			left = mid + 1
		} else {
			result = mid
			right = mid - 1
		}
	}
	return result
}

func searchRight(nums []int, target int) int {
	result := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		diff := target - nums[mid]
		if diff < 0 {
			right = mid - 1
		} else if diff > 0 {
			left = mid + 1
		} else {
			result = mid
			left = mid + 1
		}
	}
	return result
}
