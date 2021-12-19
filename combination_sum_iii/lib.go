package combinationsumiii

var result [][]int
var path []int
var pathSum int

func init() {
	result = [][]int{}
	path = []int{}
	pathSum = 0
}

func combinationSum3(k int, n int) [][]int {
	backtracking(k, n, 1)
	return result
}

func backtracking(k, n, start int) {
	if k == len(path) {
		if pathSum != n {
			return
		}
		res := []int{}
		result = append(result, append(res, path...))
		return
	}
	for i := start; i <= 9; i++ {
		if pathSum+(k-len(path))*i > n {
			continue
		}
		path = append(path, i) // push element
		pathSum += i
		backtracking(k, n, i+1)
		path = path[:len(path)-1] // pop element
		pathSum -= i
	}
}
