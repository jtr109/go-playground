// https://leetcode.com/problems/remove-element/

package removeelement

func removeElement(nums []int, val int) int {
	// 如果只是为了得到「不等于 val 的值的总数」其实用一次 O(n) 的简单查询就可以了
	// 但是为了实现修改 nums，需要做一些额外操作
	length := len(nums)
	index := 0
	for index < length {
		if nums[index] != val {
			index++
			continue
		}
		// 这里直接覆盖 index 位置的值即可，length-1 位置的值是否修改没有意义
		nums[index] = nums[length-1]
		length--
	}
	return length
}
