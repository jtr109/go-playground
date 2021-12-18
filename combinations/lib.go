// https://leetcode.com/problems/combinations/

package combinations

func combine(n int, k int) [][]int {
	s := []int{}
	for i := 1; i <= n; i++ {
		s = append(s, i)
	}
	return combineSlice(k, s)
}

func combineSlice(k int, s []int) [][]int {
	res := [][]int{}
	if k == 1 {
		for _, e := range s {
			res = append(res, []int{e})
		}
		return res
	}
	for i, e := range s {
		subRes := combineSlice(k-1, s[i+1:])
		for _, r := range subRes {
			r_ := []int{e}
			r_ = append(r_, r...)
			res = append(res, r_)
		}
	}
	return res
}
