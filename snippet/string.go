package snippet

import (
	"strings"
)

func ContainsStr(s []string, sub string) bool {
	for i := 0; i < len(s); i++ {
		if !strings.Contains(s[i], sub) {
			continue
		}
		return true
	}
	return false
}
