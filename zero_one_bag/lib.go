package zeroonebag

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func getMaxValue(weight []int, value []int, totalWeight int) int {
	totalValue := [][]int{}
	for i := 0; i < len(weight); i++ {
		totalValue = append(totalValue, []int{})
		packageWeight := weight[i]
		packageValue := value[i]
		for w := 0; w <= totalWeight; w++ {
			if packageWeight > w {
				totalValue[i] = append(totalValue[i], 0)
				continue
			}
			// max value between using current package or not
			valueWithoutPackage := 0
			if i != 0 {
				valueWithoutPackage = totalValue[i-1][w]
			}
			valueWithPackage := packageValue
			if i != 0 {
				valueWithPackage += totalValue[i-1][w-packageWeight]
			}
			totalValue[i] = append(totalValue[i], max(valueWithPackage, valueWithoutPackage))
		}
	}
	return totalValue[len(weight)-1][totalWeight]
}
