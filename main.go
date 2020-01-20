package main

import (
	"fmt"
	"github.com/go-numb/atCoder/snippet"
)

func main() {
	var nWords int = 3
	n, s, err := snippet.Scanf(nWords)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("length: %d, inputs: %s\n", n, s)
}
