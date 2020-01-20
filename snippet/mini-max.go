package snippet

import "math"

func MinInt(n ...int) int {
	if len(n) == 0 {
		panic("funciton MinInt() requires at least one argument.")
	}
	less := n[0]
	for i := 0; i < len(n); i++ {
		less = int(math.Min(float64(less), float64(n[i])))
	}
	return less
}

func MaxInt(n ...int) int {
	if len(n) == 0 {
		panic("funciton MaxInt() requires at least one argument.")
	}
	great := n[0]
	for i := 0; i < len(n); i++ {
		great = int(math.Max(float64(great), float64(n[i])))
	}
	return great
}
