package str

import "strings"

// g flag, true - ignore registr, false - not
func Contains(text, target string, g bool) bool {
	target = strings.TrimSpace(target)
	if g {
		text = strings.ToLower(text)
		target = strings.ToLower(target)
	}
	return strings.Contains(text, target)
}
