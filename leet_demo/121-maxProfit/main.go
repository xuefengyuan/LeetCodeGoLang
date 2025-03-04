package main

import "fmt"

// 买卖股票的最佳时机
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
//
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
func main() {
	//
	prices := []int{7, 1, 5, 3, 6, 4}
	result := maxProfit(prices)
	fmt.Println(result)

	prices = []int{7, 6, 4, 3, 1}
	maxProfit(prices)
	result = maxProfit(prices)
	fmt.Println(result)

	prices = []int{2, 4, 1}
	result = maxProfit(prices)
	fmt.Println(result)
}

func maxProfit(prices []int) int {

	// 初始化为第一个价格
	minPirce := prices[0]
	// 初始化最大利润为0
	maxPrice := 0
	for _, price := range prices {

		// 当前价格小于最低价格就更新最低价格
		if price < minPirce {
			minPirce = price
		}

		// 计算利润=当前价格-最低价格
		count := price - minPirce
		// 当前利润大于最大利润就更新最大利润
		if count > maxPrice {
			maxPrice = count
		}
	}
	// 返回最大利润
	return maxPrice
}
