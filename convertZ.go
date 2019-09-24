package main

import "bytes"

func convertZ1(s string, rowNum int) string {
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

func convertZ2(s string, rowNum int) string {
	if rowNum == 1 {
		return s
	}
	buf := make([]bytes.Buffer, rowNum)
	i := 0
	delta := 1
	for _, r := range s {
		buf[i].WriteRune(r)
		i += delta
		if i == rowNum-1 {
			delta = -1
		} else if i == 0 {
			delta = 1
		}
	}
	ret := bytes.Buffer{}
	for i := 0; i < rowNum; i++ {
		for _, r := range buf[i].Bytes() {
			ret.WriteByte(r)
		}
	}
	return ret.String()
}

func convert1(s string, rowNum int) string {
	sz := len(s)
	buf := make([]byte, sz)
	ref := rowNum - 1
	c := 0
	for i := 0; i < sz; i++ {
		buf[i] = s[c]
		c = 2*ref - c
		if c > sz {

		}
	}
	return ""
}
