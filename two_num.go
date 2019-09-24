package main

import "sort"

type SortedArrayIndex struct {
	data []int
	idx  []int
}

type indexNode struct {
	idx        int
	prev, next *indexNode
}

func newSortedArrayIndex(data []int) *SortedArrayIndex {
	s := &SortedArrayIndex{
		data: data,
		idx:  make([]int, len(data)),
	}
	for i := 0; i < len(data); i++ {
		s.idx[i] = i
	}
	sort.Sort(s)
	return s
}

func (s *SortedArrayIndex) Len() int {
	return len(s.data)
}

func (s *SortedArrayIndex) Less(i int, j int) bool {
	a := s.data[s.idx[i]]
	b := s.data[s.idx[j]]
	return a < b
}

func (s *SortedArrayIndex) Swap(i int, j int) {
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

func twoSum1(nums []int, target int) []int {
	s := newSortedArrayIndex(nums[:])
	a := make([]int, len(nums))
	for i, idx := range s.idx {
		n := s.data[idx]
		for sortedIdx, m := range a {
			tmp := n + m
			if tmp == target {
				if x := s.idx[sortedIdx]; x > idx {
					return []int{idx, x}
				} else {
					return []int{x, idx}
				}
			}
		}
		a[i] = n
	}
	return nil
}

func twoSum2(nums []int, target int) []int {
	a := make([]int, len(nums))
	var m int
	for i, n := range nums {
		for j := 0; j < i; j++ {
			m = a[j]
			if n+m == target {
				return []int{j, i}
			}
		}
		a[i] = n
	}
	return nil
}

// O(n) = (n^2-n)/2
