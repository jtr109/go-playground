// https://leetcode.com/problems/fruit-into-baskets/

package fruitintobackets

func totalFruit(fruits []int) int {
	// 子集合问题，使用滑动窗口模式解决

	picked := make(map[int]int)

	// 左光标移动的条件是右光标右侧的数字不在子集中，移动时要使左光标移动前坐在的值从子集中完全剔除
	// 右光标移动的条件是右光标右侧的数字在子集中，移动时要移动到右侧不是子集中已存在的数字为止
	left, right := 0, 0
	picked[fruits[left]]++
	maxPickedCount := 0
	for {
		if len(picked) >= 3 { // actually the max count should be 3
			picked[fruits[left]]--
			if picked[fruits[left]] == 0 {
				delete(picked, fruits[left])
			}
			left++
			continue
		}
		// picked is valid
		pickedCount := 0
		for _, v := range picked {
			pickedCount += v
		}
		if pickedCount > maxPickedCount {
			maxPickedCount = pickedCount
		}
		// move right cursor
		if right == len(fruits)-1 {
			// touch right boundary
			break
		}
		right++
		picked[fruits[right]]++
	}
	return maxPickedCount
}

// time complexity: O(N)
// space complexity: O(1)
