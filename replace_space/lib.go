// https://github.com/lifei6671/interview-go/blob/master/question/q005.md

package replacespace

import "strings"

func replaceSpace(s string) string {
	return strings.ReplaceAll(s, " ", "%20")
}
