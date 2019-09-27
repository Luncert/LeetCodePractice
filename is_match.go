package main

func isMatch(s string, exp string) bool {
	sz := len(s)
	expSz := len(exp)
	var idx, expIdx int
	matched := make([]*Match, 8)
	for expIdx < expSz {
		switch exp[expIdx] {
		case '*':
			pre := exp[expIdx-1] // won't check whether * is the first ch, like: '*ab'
			start := idx
			for ; idx < sz && s[idx] == pre; idx++ {
			}
			matched = append(matched, newMultiMatch(start, idx-1))
			expIdx++
		case '.':
			idx++
			expIdx++
		default:
			if s[idx] == exp[expIdx] {
				matched = append(matched, newRuneMatch(idx))
				idx++
				expIdx++
			} else if s[idx] == s[idx-1] &&
				matched[len(matched)-1].kind == kindMultiMatch {

			}
		}
	}
}

const (
	kindMultiMatch = iota
	kindSingleMatch
	kindRuneMatch
)

type Match struct {
	kind       int
	start, end int
}

func newMultiMatch(start, end int) *Match {
	return &Match{
		kind:  kindMultiMatch,
		start: start,
		end:   end,
	}
}

func newSingleMatch(pos int) *Match {
	return &Match{
		kind:  kindSingleMatch,
		start: pos,
		end:   pos,
	}
}

func newRuneMatch(pos int) *Match {
	return &Match{
		kind:  kindRuneMatch,
		start: pos,
		end:   pos,
	}
}
