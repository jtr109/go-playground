package minimumsizesubarraysum

func minSubArrayLen(target int, nums []int) int {
	// 使用双指针实现滑动窗口，目标就是两个指针不能走回头路
	// 如果子集合之和小于目标值了，右光标右移，扩大子集合
	// 如果子集合之和大于目标值了，左光标右移，缩小子集合
	minLength := 0
	i := 0
	sum := 0
	for j := 0; j < len(nums); j++ {
		sum = sum + nums[j]
		for sum >= target {
			length := j - i + 1
			if length < minLength || minLength == 0 {
				minLength = length
			}
			sum = sum - nums[i]
			i++
		}
	}
	return minLength
}
