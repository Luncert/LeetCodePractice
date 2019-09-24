package main

import "math"

func reverseInt32(x int) int {
	var ret int64
	i := 0
	buf := make([]int64, 32)
	sign := x>>32 | 1
	x *= sign
	for ; x != 0; i++ {
		buf[i] = int64(x % 10)
		x /= 10
	}
	i--
	for bitNum := i; i >= 0; i-- {
		ret += buf[i] * int64(math.Pow(10, float64(bitNum-i)))
	}
	if ret > math.MaxInt32 {
		// input is int32, so it not possible that ret becomes negative because of overflow.
		return 0
	} else {
		return int(ret) * sign
	}
}
