package notverysecure

func alphanumeric(str string) bool {
	for _, c := range str {
		if !(c >= 'A' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}
