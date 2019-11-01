package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	buf := make([]byte, 33)
	bufSz := 0
	for x != 0 {
		buf[bufSz] = byte(x % 10)
		bufSz++
		x /= 10
	}
	var (
		l, r int
	)
	if bufSz%2 == 0 {
		r = bufSz / 2
		l = r - 1
	} else {
		l = bufSz/2 - 1
		r = l + 2
	}
	for l >= 0 && r < bufSz {
		if buf[l] != buf[r] {
			return false
		}
		l--
		r++
	}
	return true
}

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	var part, b int
	b = x % 10
	// e.g. abc0, i return false because a won't be a
	if b == 0 {
		return false
	}
	x /= 10
	for x != 0 {
		if x == part {
			return true
		}
		part = part*10 + b
		if x == part {
			return true
		}

		b = x % 10
		x /= 10
	}
	return false
}
