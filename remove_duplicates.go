package main

func removeDuplicates(nums []int) int {
	sz := len(nums)
	if sz <= 1 {
		return sz
	}
	idx := 1
	prevIdx := 0
	var step int
	for idx < sz {
		if nums[idx] == nums[prevIdx] {
			step++
		} else {
			if step > 0 {
				nums = append(nums[:prevIdx], nums[prevIdx+step:]...)
				idx -= step
				sz -= step
				step = 0
			}
			prevIdx = idx
		}
		idx++
	}
	if step > 0 {
		nums = append(nums[:prevIdx], nums[prevIdx+step:]...)
		sz -= step
	}
	return sz
}

func removeDuplicatesCount(nums []int) int {
	sz := len(nums)
	if sz <= 1 {
		return sz
	}
	ret := 1
	idx := 1
	prevIdx := 0
	for idx < sz {
		if nums[idx] != nums[prevIdx] {
			ret++
			prevIdx = idx
		}
		idx++
	}
	return ret
}
