package main

func rotate1(nums []int, k int) {
	sz := len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[sz-1]
		for j := sz - 1; j > 0; j-- {
			nums[j] = nums[j-1]
		}
		nums[0] = tmp
	}
}

func rotate2(nums []int, k int) {
	sz := len(nums)
	if k == 0 {
		return
	} else if k > sz {
		k = k % sz
	}
	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = nums[sz-k+i]
	}
	for i := sz - 1; i > k-1; i-- {
		nums[i] = nums[i-k]
	}
	for i, v := range tmp {
		nums[i] = v
	}
}
