// https://leetcode.com/problems/move-zeroes/

package movezeros

func moveZeroes(nums []int) {
	// 可以维护两个指针，一个记录待替换元素的索引，一个记录正在比较的元素的索引
	slowIndex := 0
	fastIndex := 0
	for fastIndex < len(nums) {
		if nums[fastIndex] == 0 {
			fastIndex++
			continue
		}
		nums[slowIndex], nums[fastIndex] = nums[fastIndex], nums[slowIndex]
		slowIndex++
		fastIndex++
	}
}

// time complexity: O(n)
// space complexity: O(1)
