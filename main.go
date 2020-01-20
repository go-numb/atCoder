package main

import (
	"fmt"
	"github.com/go-numb/atCoder/snippet"
)

func main() {
	var nWords int = 3
	n, s := snippet.Scanf(nWords)
	if n == 0 {
		fmt.Println("has not data")
	}
	fmt.Printf("length: %d, inputs: %s\n", n, s)
}
