// https://leetcode.com/explore/challenge/card/september-leetcoding-challenge-2021/637/week-2-september-8th-september-14th/3973/

package maximumnumberofballoons

import (
	"strings"
)

func maxNumberOfBalloons(text string) int {
	chars := []rune{'b', 'a', 'l', 'o', 'n'}
	charCounts := make(map[rune]int)
	for _, c := range chars {
		charCounts[c] = strings.Count(text, string(c))
	}
	maxWordCounts := []int{}
	for c, count := range charCounts {
		if strings.ContainsRune("lo", c) {
			maxWordCounts = append(maxWordCounts, count/2)
		} else {
			maxWordCounts = append(maxWordCounts, count)
		}
	}
	max := maxWordCounts[0]
	for _, count := range maxWordCounts {
		if count < max {
			max = count
		}
	}
	return max
}
