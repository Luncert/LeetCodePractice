package main

import "sort"

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, v := range nums {
		m[v]++
		if m[v] == 2 {
			return true
		}
	}
	return false
}

type IntSlice []int

func (is IntSlice) Len() int {
	return len(is)
}

func (is IntSlice) Less(i, j int) bool {
	return is[i] < is[j]
}

func (is IntSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func containsDuplicate1(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sort.Sort(IntSlice(nums))
	prev := nums[0] - 1
	for _, v := range nums {
		if v == prev {
			return true
		} else {
			prev = v
		}
	}
	return false
}
