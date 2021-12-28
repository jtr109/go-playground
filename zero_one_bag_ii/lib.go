package zeroonebagii

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMaxValue(weight []int, value []int, targetWeight int) int {
	dp := []int{}
	for w := 0; w <= targetWeight; w++ {
		dp = append(dp, 0)
	}
	for i := 0; i < len(weight); i++ {
		packageWeight := weight[i]
		packageValue := value[i]
		for w := targetWeight; w >= 0; w-- {
			if w < packageWeight {
				continue
			}
			dp[w] = max(dp[w], packageValue+dp[w-packageWeight])
		}
	}
	return dp[targetWeight]
}
