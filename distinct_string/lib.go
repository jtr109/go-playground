// https://github.com/lifei6671/interview-go/blob/1576d64ec1bdd311e8ad7e1132217c02d708b38e/question/q002.md

package distinctstring

import "strings"

func isDistinct(s string) bool {
	for i, c := range s {
		if strings.Index(s, string(c)) != i {
			return false
		}
	}
	return true
}
