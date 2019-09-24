package main

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
