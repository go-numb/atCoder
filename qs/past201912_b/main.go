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
			rank, diff := ranking(sale, sales)
			fmt.Printf("today sale rank%d and diff%d in Sales\n", rank, diff)
		}
		sales = append(sales, sale)
	}
}

func ranking(n int, num []int) (rank, diff int) {
	diff = n - num[len(num)-1]
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Printf("%+v\n", num)
	for i := range num {
		if num[i] <= n {
			fmt.Printf("%d: %d -> %d\n", i, n, num[i])
			return i + 1, diff
		}
	}
	// 最下位
	return len(num) + 1, diff
}
