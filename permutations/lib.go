package permutations

func permute(nums []int) (result [][]int) {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}
	permuted := permute(nums[1:])
	for _, perm := range permuted {
		for i := 0; i <= len(perm); i++ {
			// insert nums[0] into all split in the slice perm
			r := []int{}
			r = append(r, perm[:i]...)
			r = append(r, nums[0])
			r = append(r, perm[i:]...)
			result = append(result, r)
		}
	}
	return result
}
