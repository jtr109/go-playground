package largestplussign

func orderOfLargestPlusSign(n int, mines [][]int) int {
	grid := createGrid(n, mines)
	left := initializeLargestOfDirectionArray(n)
	right := initializeLargestOfDirectionArray(n)
	up := initializeLargestOfDirectionArray(n)
	down := initializeLargestOfDirectionArray(n)
	order := 0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if x := getLargestPlusSignOfCurrentElement(grid, left, right, up, down, r, c, n); x > order {
				order = x
			}
		}
	}
	return order
}

func createGrid(n int, mines [][]int) (grid [][]int) {
	for r := 0; r < n; r++ {
		row := []int{}
		for c := 0; c < n; c++ {
			row = append(row, 1)
		}
		grid = append(grid, row)
	}
	for _, m := range mines {
		r := m[0]
		c := m[1]
		grid[r][c] = 0
	}
	return
}

func initializeLargestOfDirectionArray(n int) (result [][]int) {
	for r := 0; r < n; r++ {
		row := []int{}
		for c := 0; c < n; c++ {
			row = append(row, -1)
		}
		result = append(result, row)
	}
	return
}

func getLargestLeftOfCurrentElement(grid [][]int, left [][]int, r int, c int, n int) int {
	if left[r][c] != -1 {
		return left[r][c]
	}
	var res int
	if grid[r][c] == 0 {
		res = 0
	} else if c == 0 {
		res = 1
	} else {
		res = getLargestLeftOfCurrentElement(grid, left, r, c-1, n) + 1
	}
	left[r][c] = res
	return res
}

func getLargestRightOfCurrentElement(grid [][]int, left [][]int, r int, c int, n int) int {
	if left[r][c] != -1 {
		return left[r][c]
	}
	var res int
	if grid[r][c] == 0 {
		res = 0
	} else if c == n-1 {
		res = 1
	} else {
		res = getLargestRightOfCurrentElement(grid, left, r, c+1, n) + 1
	}
	left[r][c] = res
	return res
}

func getLargestUpOfCurrentElement(grid [][]int, left [][]int, r int, c int, n int) int {
	if left[r][c] != -1 {
		return left[r][c]
	}
	var res int
	if grid[r][c] == 0 {
		res = 0
	} else if r == 0 {
		res = 1
	} else {
		res = getLargestUpOfCurrentElement(grid, left, r-1, c, n) + 1
	}
	left[r][c] = res
	return res
}

func getLargestDownOfCurrentElement(grid [][]int, left [][]int, r int, c int, n int) int {
	if left[r][c] != -1 {
		return left[r][c]
	}
	var res int
	if grid[r][c] == 0 {
		res = 0
	} else if r == n-1 {
		res = 1
	} else {
		res = getLargestDownOfCurrentElement(grid, left, r+1, c, n) + 1
	}
	left[r][c] = res
	return res
}

func getLargestPlusSignOfCurrentElement(grid [][]int, left [][]int, right [][]int, up [][]int, down [][]int, r int, c int, n int) int {
	largestPlus := getLargestLeftOfCurrentElement(grid, left, r, c, n)
	if x := getLargestRightOfCurrentElement(grid, right, r, c, n); x < largestPlus {
		largestPlus = x
	}
	if x := getLargestUpOfCurrentElement(grid, up, r, c, n); x < largestPlus {
		largestPlus = x
	}
	if x := getLargestDownOfCurrentElement(grid, down, r, c, n); x < largestPlus {
		largestPlus = x
	}
	return largestPlus
}
