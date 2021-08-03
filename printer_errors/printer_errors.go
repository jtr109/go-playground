package printererrors

import (
	"fmt"
	"strings"
)

func PrinterError(s string) string {
	errorCount := len(s)
	for r := 'a'; r <= 'm'; r++ {
		errorCount -= strings.Count(s, string(r))
	}
	return fmt.Sprintf("%d/%d", errorCount, len(s))
}
