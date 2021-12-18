package combinationsumiii

func combinationSum3(k int, n int) [][]int {
	s := []int{}
	for i := 1; i <= 9; i++ {
		s = append(s, i)
	}
	return combine(k, n, s)
}

func combine(k, n int, s []int) [][]int {
	if k > len(s) {
		return [][]int{}
	}
	if k == 1 {
		if n >= s[0] && n <= s[len(s)-1] {
			return [][]int{{n}}
		} else {
			return [][]int{}
		}
	}
	res := [][]int{}
	for i, m := range s {
		for _, r := range combine(k-1, n-m, s[i+1:]) {
			res = append(res, append([]int{m}, r...))
		}
	}
	return res
}
