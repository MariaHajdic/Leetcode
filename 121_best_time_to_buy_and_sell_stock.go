package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for i := range prices {
		maxProfit = max(prices[i]-minPrice, maxProfit)
		minPrice = min(minPrice, prices[i])
	}

	return maxProfit
}
