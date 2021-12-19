// https://leetcode.com/problems/combinations/

package combinations

var result [][]int
var path []int

func init() {
	result = [][]int{}
	path = []int{}
}

func backtracking(n, k, start int) {
	if k == len(path) {
		r := []int{}
		result = append(result, append(r, path...))
		return
	}
	for i := start; i <= n-(k-len(path))+1; i++ {
		path = append(path, i)
		backtracking(n, k, i+1)
		path = path[:len(path)-1]
	}
}

func combine(n, k int) [][]int {
	backtracking(n, k, 1)
	return result
}
