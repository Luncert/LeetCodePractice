package main

func maxProfit(prices []int) int {
	sz := len(prices)
	if sz <= 1 {
		return 0
	}
	profit := 0
	prev := prices[0]
	for i := 1; i < sz; i++ {
		if prices[i] > prev {
			profit += prices[i] - prev
		}
		prev = prices[i]
	}
	return profit
}
