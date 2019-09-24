package main

func lengthOfLongestSubstring(s string) (ret int) {
	ret = 0
	sz := len(s)
	if sz <= 1 {
		return sz
	}

	var r uint8
	ri := newRuneIndex()
	i := 0
	start := 0
	var tmp int

	for true {
		r = s[i]
		if pos, ok := ri.find(r); ok {
			tmp = i - start
			if ret < tmp {
				ret = tmp
			}
			tmp = pos + 1
			if start < tmp {
				start = tmp
			}
		}
		ri.set(r, i)

		i++
		if i == sz {
			tmp := i - start
			if ret < tmp {
				ret = tmp
			}
			break
		}
	}
	return
}

type RuneIndex struct {
	index []int
	flag  []bool
}

func newRuneIndex() *RuneIndex {
	return &RuneIndex{
		index: make([]int, 256),
		flag:  make([]bool, 256),
	}
}

func (ri *RuneIndex) set(r uint8, pos int) {
	ri.flag[r] = true
	ri.index[r] = pos
}

func (ri *RuneIndex) find(r uint8) (p int, ok bool) {
	if ri.flag[r] {
		p = ri.index[r]
		ok = true
	}
	return
}

//func (ri *RuneIndex) parseRuneToIndex(r uint8) int {
//	if r >= 'a' && r <= 'z' {
//		return int(r - 'a')
//	} else {
//		return int(r - 'A')
//	}
//}
