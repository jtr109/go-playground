package consecutivestrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func dotest(t *testing.T, strarr []string, k int, expected string) {
	assert.Equal(t, expected, LongestConsec(strarr, k))
}

func TestSample(t *testing.T) {
	// cSpell: disable
	dotest(t, []string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta")
	dotest(t, []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
		"oocccffuucccjjjkkkjyyyeehh")
	dotest(t, []string{}, 3, "")
	dotest(t, []string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
		"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck")
	// cSpell: enable
}

func TestSubmission1(t *testing.T) {
	// cSpell: disable
	dotest(t, []string{"nnbbsssv", "wwwaavfffw", "mmmqaaappplfffw", "qqwffccc", "llouuaa", "iiiqfnhuujj", "bffwwx", "yyllwwoo", "gghsss"}, 3, "nnbbsssvwwwaavfffwmmmqaaappplfffw")
	// cSpell: enable
}
