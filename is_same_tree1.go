package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const minInt = ^int(uint(0) >> 1)

func isSameTree1(p *TreeNode, q *TreeNode) bool {
	sp := make([]*TreeNode, 0)
	sq := make([]*TreeNode, 0)
	sp = append(sp, p)
	sq = append(sq, q)
	for len(sp) > 0 && len(sq) > 0 {
		n1 := sp[0]
		sp = sp[1:]
		n2 := sq[0]
		sq = sq[1:]
		a := n1 == nil
		b := n2 == nil
		if a && b {
			continue
		} else if a || b || n1.Val != n2.Val {
			return false
		}
		sp = append(sp, n1.Left, n1.Right)
		sq = append(sq, n2.Left, n2.Right)
	}
	return len(sq) == 0
}
