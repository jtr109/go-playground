// https://leetcode.com/problems/unique-paths-ii/

package uniquepathsii

func newCache(obstacleGrid [][]int) [][]int {
	cache := [][]int{}
	for i := 0; i < len(obstacleGrid); i++ {
		cache = append(cache, []int{})
		for j := 0; j < len(obstacleGrid[i]); j++ {
			cache[i] = append(cache[i], -1)
		}
	}
	return cache
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	cache := newCache(obstacleGrid)
	return dp(0, 0, obstacleGrid, cache)
}

func dp(i, j int, obstacleGrid, cache [][]int) int {
	if cache[i][j] != -1 {
		return cache[i][j]
	}
	if obstacleGrid[i][j] == 1 {
		cache[i][j] = 0
		return 0
	}
	atBottom := i == len(obstacleGrid)-1
	atRight := j == len(obstacleGrid[i])-1
	if atBottom && atRight {
		res := 1
		if cache[i][j] == 1 {
			res = 0
		}
		cache[i][j] = res
		return res
	}
	result := 0
	if !atBottom {
		result += dp(i+1, j, obstacleGrid, cache)
	}
	if !atRight {
		result += dp(i, j+1, obstacleGrid, cache)
	}
	cache[i][j] = result
	return result
}
