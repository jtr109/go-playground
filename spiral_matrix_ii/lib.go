package spiralmatrixii

const (
	RIGHT = iota
	DOWN
	LEFT
	UP
)

func generateMatrix(n int) [][]int {
	matrix := initializeMatrix(n)
	i := 0
	j := 0
	number := 1
	for lineLength := n - 1; lineLength > 1; lineLength-- {
		// right move
		for step := 0; step < lineLength; step++ {
			matrix[i][j] = number
			number++
			j++
		}
		// down move
		for step := 0; step < lineLength; step++ {
			matrix[i][j] = number
			number++
			i++
		}
		// left move
		for step := 0; step < lineLength; step++ {
			matrix[i][j] = number
			number++
			j--
		}
		// up move
		for step := 0; step < lineLength-1; step++ {
			matrix[i][j] = number
			number++
			i--
		}
		matrix[i][j] = number
		number++
		j++
		matrix[i][j] = number
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

func changeDirection(direction int) (newDirection int) {
	switch direction {
	case RIGHT:
		newDirection = DOWN
	case DOWN:
		newDirection = LEFT
	case LEFT:
		newDirection = UP
	case UP:
		newDirection = RIGHT
	}
	return
}
