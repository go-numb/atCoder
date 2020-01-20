package snippet

import (
	"fmt"
	"strings"
)

func Scanf(n int) (count int, s string) {
	str := make([]string, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&s)
		if err != nil {
			continue
		}
		count += len([]rune(s))
		str[i] = s
	}

	return count + n - 1, strings.Join(str, " ")
}
