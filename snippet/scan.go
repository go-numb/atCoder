package snippet

import (
	"fmt"
	"strings"
)

func Scanf(n int) (int, string, error) {
	var totalLength int
	str := make([]string, n)
	for i := 0; i < n; i++ {
		var s string
		length, err := fmt.Scan(&s)
		if err != nil {
			continue
		}
		fmt.Printf("%d - %+v - %s\n", length, len([]rune(s)), s)
		totalLength += length
		str[i] = s
	}

	return totalLength + n - 1, strings.Join(str, " "), nil
}
