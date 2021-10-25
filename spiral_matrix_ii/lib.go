// https://leetcode.com/problems/spiral-matrix-ii/

package spiralmatrixii

func generateMatrix(n int) [][]int {
	matrix := initializeMatrix(n)
	i := 0
	j := 0
	number := 1
	for ; n > 0; n = n - 2 {
		matrix[i][j] = number
		if n == 1 {
			break
		}
		// right move
		for step := 0; step < n-1; step++ {
			number++
			j++
			matrix[i][j] = number
		}
		// down move
		for step := 0; step < n-1; step++ {
			number++
			i++
			matrix[i][j] = number
		}
		// left move
		for step := 0; step < n-1; step++ {
			number++
			j--
			matrix[i][j] = number
		}
		// up move
		for step := 0; step < n-2; step++ {
			number++
			i--
			matrix[i][j] = number
		}
		// move into the inner circle
		number++
		j++
	}
	return matrix
}

func initializeMatrix(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		line := []int{}
		for j := 0; j < n; j++ {
			line = append(line, 0)
		}
		result = append(result, line)
	}
	return result
}
