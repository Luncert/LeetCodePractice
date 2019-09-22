package main

import (
	"bytes"
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (ret *ListNode) {
	var carry int
	if l1 != nil && l2 != nil {
		ret, carry = calc(carry, l1.Val, l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	} else {
		ret = &ListNode{}
	}
	node := ret
	for l1 != nil && l2 != nil {
		node.Next, carry = calc(carry, l1.Val, l2.Val)
		node = node.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	y := func(l *ListNode) {
		for ; l != nil; l = l.Next {
			node.Next, carry = calc(carry, l.Val)
			node = node.Next
		}
	}
	y(l1)
	y(l2)
	if carry != 0 {
		node.Next = &ListNode{Val: carry}
	}
	return
}

func calc(preCarry int, n ...int) (l *ListNode, carry int) {
	l = &ListNode{}
	for _, i := range n {
		l.Val += i
	}
	l.Val += preCarry
	if l.Val >= 10 {
		l.Val %= 10
		carry = 1
	}
	return
}

func buildList(n ...int) *ListNode {
	ret := &ListNode{Val: n[0]}
	node := ret
	for i := 1; i < len(n); i++ {
		node.Next = &ListNode{Val: n[i]}
		node = node.Next
	}
	return ret
}

func printList(l *ListNode) {
	buf := bytes.Buffer{}
	for l != nil {
		buf.WriteString(strconv.FormatInt(int64(l.Val), 10))
		buf.WriteRune(' ')
		l = l.Next
	}
	fmt.Println(buf.String())
}
