package main

import "fmt"

const minInt = ^int(^uint(0) >> 1)

func testFindMaxSubArray(prices []int) int {
	sz := len(prices) - 1
	for i := 0; i < sz; i++ {
		prices[i] = prices[i+1] - prices[i]
	}
	prices[sz] = 0
	s, e, profit := findMaxSubArray(prices, 0, sz-1)
	fmt.Println(prices, s, e)
	return profit
}

func findMaxSubArray(data []int, low, high int) (int, int, int) {
	if low == high {
		return low, high, data[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := findMaxSubArray(data, low, mid)
		rightLow, rightHigh, rightSum := findMaxSubArray(data, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubArray(data, low, mid, high)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

func findMaxCrossingSubArray(data []int, low, mid, high int) (int, int, int) {
	var (
		maxLeft, maxRight int
		leftSum           = minInt
		rightSum          = minInt
		sum               = 0
	)
	for i := mid; i >= low; i-- {
		sum += data[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	sum = 0
	for i := mid + 1; i <= high; i++ {
		sum += data[i]
		if sum > rightSum {
			rightSum = sum
			maxRight = i
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}
