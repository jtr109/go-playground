package flipstringtomonotoneincreasing

import "strings"

func minFlipsMonoIncr(s string) int {
	flipCount := []int{strings.Count(s, "0")} // flip all into 1
	for _, c := range s {
		if c == '0' {
			flipCount = append(flipCount, flipCount[len(flipCount)-1]-1)
		} else {
			flipCount = append(flipCount, flipCount[len(flipCount)-1]+1)
		}
	}
	minCount := flipCount[0]
	for _, c := range flipCount {
		if c < minCount {
			minCount = c
		}
	}
	return minCount
}
