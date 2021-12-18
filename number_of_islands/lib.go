package numberofislands

// func numIslands(grid [][]byte) int {
// 	count := 0
// 	for i := 0; i < len(grid); i++ {
// 		for j := 0; j < len(grid[i]); j++ {
// 			if grid[i][j] == 0 {
// 				continue
// 			}
// 			if (i == 0 || grid[i-1][j] == 0) && (j == 0 || grid[i][j-1] == 0) {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				continue
			} else {
				count++
				dfs(grid, i, j)
			}

		}
	}
	return count
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i == len(grid) || j == len(grid[i]) {
		return
	}
	if grid[i][j] == 0 {
		return
	}
	if grid[i][j] == 1 {
		grid[i][j] = 0
	}
	dfs(grid, i-1, j)
	dfs(grid, i, j-1)
	dfs(grid, i+1, j)
	dfs(grid, i, j+1)
}
