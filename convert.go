package main

import "bytes"

func convert(s string, rowNum int) string {
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
