package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		sales []int
		sale  int
	)

	for {
		_, err := fmt.Scan(&sale)
		if err != nil {
			fmt.Println("input undefined, not number")
			continue
		}
		if 0 < len(sales) {
			fmt.Printf("today sale rank %d in Sales\n", ranking(sale, sales))
		}
		sales = append(sales, sale)
	}
}

func ranking(n int, num []int) int {
	sort.Ints(num)
	for i := range num {
		if num[i] < n {
			return i + 1
		}
	}
	return len(num) + 1
}
