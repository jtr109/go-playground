package simplifypath

import "strings"

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	result := []string{}
	for _, p := range paths {
		if p == "." || p == "" {
			continue
		}
		if p == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
			continue
		}
		result = append(result, p)
	}
	return "/" + strings.Join(result, "/")
}
