package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		sz1        = len(nums1)
		sz2        = len(nums2)
		halfSz     = (sz1+sz2)/2 + 1
		buf        = []int{0, 0}
		i, j, k, x int
	)
	for k < halfSz {
		if i == sz1 {
			if j < sz2 {
				buf[x], j = nums2[j], j+1
			} else {
				break
			}
		} else if j == sz2 {
			buf[x], i = nums1[i], i+1
		} else if nums1[i] <= nums2[j] {
			buf[x], i = nums1[i], i+1
		} else {
			buf[x], j = nums2[j], j+1
		}
		k++
		x = (x + 1) % 2
	}
	if (sz1+sz2)%2 == 0 {
		return float64(buf[0]+buf[1]) / 2
	}
	if x == 0 {
		return float64(buf[1])
	}
	return float64(buf[0])
}
