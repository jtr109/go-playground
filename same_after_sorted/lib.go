// https://github.com/lifei6671/interview-go/blob/master/question/q004.md

package sameaftersorted

import "strings"

func sameAfterSorted(s1 string, s2 string) bool {
	if len(s1) != len(s2) || len(s1) > 5000 {
		return false
	}
	for _, r := range s1 {
		if strings.Count(s1, string(r)) != strings.Count(s2, string(r)) {
			return false
		}
	}
	return true
}
