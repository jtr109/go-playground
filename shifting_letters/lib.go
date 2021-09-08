// https://leetcode.com/explore/challenge/card/september-leetcoding-challenge-2021/637/week-2-september-8th-september-14th/3968/

package shiftingletters

func shiftingLetters(s string, shifts []int) string {
	charShifts := getCharShifts(shifts)
	result := []rune{}
	for i, c := range s {
		result = append(result, shiftMultipleTimes(c, charShifts[i]))
	}
	return string(result)
}

func getCharShifts(shifts []int) []int {
	for i, j := 0, len(shifts)-1; i < j; i, j = i+1, j-1 {
		shifts[i], shifts[j] = shifts[j], shifts[i]
	}
	charShifts := []int{}
	last := 0
	for _, x := range shifts {
		last += x
		charShifts = append(charShifts, last)
	}
	for i, j := 0, len(charShifts)-1; i < j; i, j = i+1, j-1 {
		charShifts[i], charShifts[j] = charShifts[j], charShifts[i]
	}
	return charShifts
}

func shiftMultipleTimes(c rune, times int) rune {
	return rune((int(c)+times-int('a'))%26 + int('a'))
}
