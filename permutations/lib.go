// https://leetcode.com/problems/permutations/

package permutations

func permute(nums []int) (result [][]int) {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	for i, n := range nums {
		others := []int{}
		others = append(others, nums[:i]...)
		others = append(others, nums[i+1:]...)
		permuted := permute(others)
		for _, perm := range permuted {
			r := []int{n}
			r = append(r, perm...)
			result = append(result, r)
		}
	}
	return
}
