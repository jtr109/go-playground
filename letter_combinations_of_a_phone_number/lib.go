package lettercombinationsofaphonenumber

import (
	"strconv"
)

var result []string
var digitMap map[int][]rune

func init() {
	result = []string{}
	digitMap = make(map[int][]rune)
	digitMap[2] = []rune("abc")
	digitMap[3] = []rune("def")
	digitMap[4] = []rune("ghi")
	digitMap[5] = []rune("jkl")
	digitMap[6] = []rune("mno")
	digitMap[7] = []rune("pqrs")
	digitMap[8] = []rune("tuv")
	digitMap[9] = []rune("wxyz")
}

func digitsToInts(digits string) []int {
	result := []int{}
	for num, _ := strconv.Atoi(digits); num != 0; num /= 10 {
		result = append(result, num%10)
	}
	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}

func letterCombinations(digits string) []string {
	backtracking(digitsToInts(digits), 0, []rune{})
	return result
}

func backtracking(digits []int, index int, path []rune) {
	if index == len(digits) {
		result = append(result, string(path))
		return
	}
	for _, v := range digitMap[digits[index]] {
		path = append(path, v)
		backtracking(digits, index+1, path)
		path = path[:len(path)-1]
	}
}
