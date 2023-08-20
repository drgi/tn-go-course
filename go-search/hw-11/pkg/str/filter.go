package str

import (
	"regexp"
	"strings"
)

func FilterString(text string) string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	re := regexp.MustCompile(`[^a-zA-Z ]`)
	return re.ReplaceAllString(text, "")
}
