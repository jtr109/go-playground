package strstr

func strStr1(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	haystackRunes := []rune(haystack)
	needleRunes := []rune(needle)
	for i := 0; i < len(haystackRunes); i++ {
		j := 0
		for j < len(needleRunes) && i+j < len(haystackRunes) {
			if haystackRunes[i+j] != needleRunes[j] {
				break
			}
			j++
		}
		if j == len(needleRunes) {
			return i
		}
	}
	return -1
}

func buildNext(needle string) (next []int) {
	// next 中每个索引对应的值意味着包含当前索引对应元素的后缀可以匹配的前缀的最后一个索引
	// 如果包含当前索引的元素的后缀没有可以匹配前缀，则将值设置为 -1
	needleRunes := []rune(needle)
	j := -1                // j+1 是前缀最后一个元素的索引
	next = append(next, j) // 第一个元素作为后缀，不存在匹配的前缀
	for i := 1; i < len(needleRunes); i++ {
		for j >= 0 && needleRunes[i] != needleRunes[j+1] {
			// 预期 i 和 j+1 索引的元素相同
			// 如果 i 和 j+1 索引对应的字符串不同，则回退 j
			// 直到到索引 i 和 j+1 对应的字符相同为止
			// 或者 j 已经退到 -1 表示没有可以匹配的前缀，则停止回退
			j = next[j]
		}
		// 如果 i 和 j+1 索引元素相同，记录 i 索引匹配的前缀最后一个元素的索引为 j+1
		// 否则 j 退回到 -1，索引 i 对应的元素为 -1
		if needleRunes[i] == needleRunes[j+1] {
			j++
		}
		next = append(next, j)
	}
	return
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	next := buildNext(needle)
	haystackRunes := []rune(haystack)
	needleRunes := []rune(needle)
	j := -1
	for i := 0; i < len(haystackRunes); i++ {
		for j >= 0 && needleRunes[j+1] != haystackRunes[i] {
			j = next[j]
		}
		if needleRunes[j+1] == haystackRunes[i] {
			j++
			if j == len(needleRunes)-1 {
				return i - j
			}
			continue
		}
		continue
	}
	return -1
}
