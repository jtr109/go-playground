package subsetsii

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		for j := i; j <= len(nums); j++ {
			n := nums[i:j]
			if !contains(result, n) {
				result = append(result, n)
			}
		}
	}
	return result
}

func contains(arr [][]int, sub []int) bool {
	contain := false
	for _, a := range arr {
		if equal(a, sub) {
			contain = true
			break
		}
	}
	return contain
}

func equal(current []int, target []int) bool {
	if len(current) != len(target) {
		return false
	}
	for i, c := range current {
		if c != target[i] {
			return false
		}
	}
	return true
}
