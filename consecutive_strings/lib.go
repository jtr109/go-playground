// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/go

package consecutivestrings

import (
	"strings"
)

func LongestConsec(strarr []string, k int) string {
	if len(strarr) < k {
		return ""
	} else if len(strarr) <= k {
		return strings.Join(strarr, "")
	}

	var lenarr []int
	for _, s := range strarr {
		lenarr = append(lenarr, len(s))
	}
	var maxLength int
	for _, l := range lenarr[:k] {
		maxLength += l
	}
	maxI := 0
	length := maxLength
	for i := 1; i < len(strarr)-k+1; i++ {
		length = length - lenarr[i-1] + lenarr[i+k-1]
		if length > maxLength {
			maxLength = length
			maxI = i
		}
	}
	return strings.Join(strarr[maxI:maxI+k], "")
}
