package snippet

import "math"

func AbsInt(a int) int {
	return int(math.Abs(float64(a)))
}

func PowInt(p, q int) int {
	return int(math.Pow(float64(p), float64(q)))
}
