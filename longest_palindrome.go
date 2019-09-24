package main

func longestPalindrome(s string) string {
	sz := len(s)
	if sz == 0 {
		return ""
	} else if sz == 1 {
		return s
	}

	var (
		start, subStrSz int
		l, r            int
	)
	for i := 1; i < sz; i++ {
		// handle aba, s[i] = b, s[l] = a, s[r] = a
		r = i + 1
		l = i - 1
		for l >= 0 && r < sz && s[l] == s[r] {
			l--
			r++
		}
		if subStrSz < r-l-1 {
			start = l + 1
			subStrSz = r - l - 1
		}
		// handle abba, s[i] = b, s[l] = b, s[r] = b
		l = i - 1
		r = i
		for l >= 0 && r < sz && s[l] == s[r] {
			l--
			r++
		}
		if subStrSz < r-l-1 {
			start = l + 1
			subStrSz = r - l - 1
		}
	}
	return s[start : start+subStrSz]
}
