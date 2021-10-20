// https://leetcode.com/problems/squares-of-a-sorted-array/

package squaresofasortedarray

func sortedSquares(nums []int) []int {
	// 因为 nums 是从小到大顺序排列的，所以平方数最大的不是在最左侧就是在最右侧
	// 所以定义两个指针分别从两侧逼近中间就可以得到一个从大到小的平方数集合
	left, right := 0, len(nums)-1
	result := []int{}
	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			result = append(result, nums[left]*nums[left])
			left++
		} else {
			result = append(result, nums[right]*nums[right])
			right--
		}
	}
	reverse(result)
	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
