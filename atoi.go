package main

import "math"

func myAtoi(str string) int {
	sz := len(str)
	var i int
	var sign int = 1
	for ; i < sz; i++ {
		if str[i] != ' ' {
			break
		}
	}
	if i < sz {
		buf := make([]byte, sz)
		bufIdx := 0
		if str[i] == '+' || str[i] == '-' {
			sign = -(int(str[i]) - 44)
			i++
		}
		for ; i < sz; i++ {
			if str[i] != '0' {
				break
			}
		}
		for ; i < sz; i++ {
			if str[i] >= '0' && str[i] <= '9' {
				buf[bufIdx] = str[i] - '0'
				bufIdx++
			} else {
				break
			}
		}
		if bufIdx >= 12 {
			goto overflow
		} else if bufIdx > 0 {
			bufIdx--
			var ret int64
			for bitNum := bufIdx; bufIdx >= 0; bufIdx-- {
				ret += int64(buf[bufIdx]) * pow10(bitNum-bufIdx)
			}
			if ret > math.MaxInt32 || ret < 0 {
				goto overflow
			} else {
				return int(ret) * sign
			}
		} else {
			return 0
		}
	}
	return 0
overflow:
	if sign > 0 {
		return math.MaxInt32
	} else {
		return math.MinInt32
	}
}

// ignore overflow
func pow10(n int) (ret int64) {
	ret = 1
	for i := 0; i < n; i++ {
		ret *= 10
	}
	return ret
}
