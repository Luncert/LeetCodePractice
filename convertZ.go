package main

func convertZ(s string, rowNum int) string {
	var (
		sz                  = len(s)
		refDelta            = rowNum - 1
		ref                 = refDelta
		idx, rowPos, bufIdx int
		buf                 = make([]byte, sz)
	)
	if rowNum <= 1 {
		return s
	}
	for ; bufIdx < sz; bufIdx++ {
		buf[bufIdx] = s[idx]
		idx = 2*ref - idx
		ref += refDelta
		if idx == ref {
			ref += refDelta
		}
		if idx >= sz {
			rowPos++
			idx = rowPos
			ref = refDelta
			if rowPos == rowNum-1 {
				bufIdx++
				break
			}
		}
	}
	for ; bufIdx < sz; bufIdx++ {
		buf[bufIdx] = s[idx]
		idx += 2 * refDelta
	}
	return string(buf)
}
