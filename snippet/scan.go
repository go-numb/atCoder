package snippet

import (
	"fmt"
	"strings"
)

func Scanf(n int) (count int, s string) {
	str := make([]string, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&str[i])
		if err != nil {
			continue
		}
		count += len([]rune(str[i]))
	}

	return count + n - 1, strings.Join(str, " ")
}
