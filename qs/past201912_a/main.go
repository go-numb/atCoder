package main

import (
	"fmt"
	"github.com/go-numb/atCoder/snippet"
	"strconv"
	"strings"
)

func main() {
	n, s := snippet.Scanf(1)
	if n != 3 {
		panic(fmt.Sprintf("length: %d, has not input, number string -> %s", n, s))
	}

	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(fmt.Sprintf("length: %d, get error, input string to number, number string -> %s", n, s))
	}

	fmt.Println(num * 2)
}
