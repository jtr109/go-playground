// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

package removeduplicatesfromsortedarray

func removeDuplicates(nums []int) int {
	// 因为操作后的 nums 要求各元素保持相对位置，所以不能使用末端元素替换当前索引位置的元素
	// 这时候需要使用双指针的方式进行操作
	slowIndex := 0
	fastIndex := 0
	for fastIndex < len(nums) {
		if fastIndex > 0 && nums[fastIndex] == nums[fastIndex-1] {
			fastIndex++
			continue
		}
		nums[slowIndex] = nums[fastIndex]
		slowIndex++
		fastIndex++
	}
	return slowIndex
}
