package main

import "fmt"

// 122.买卖股票最佳时机II
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
//
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
//
// 返回 你能获得的 最大 利润
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	ret := maxProfit(prices)
	fmt.Println(ret)

	prices = []int{1, 2, 3, 4, 5}
	ret = maxProfit(prices)
	fmt.Println(ret)

	prices = []int{7, 6, 4, 3, 1}
	ret = maxProfit(prices)
	fmt.Println(ret)
}

func maxProfit(prices []int) int {

	// 初始化利润
	count := 0
	// 开始遍历，注意下标问题
	for i := 0; i < len(prices)-1; i++ {
		// 判断，当前下标后一个元素比较当前元素大
		if prices[i+1] > prices[i] {
			// 下一个元素减当前元素，利润累加
			count += prices[i+1] - prices[i]
		}
	}
	return count
}

func maxProfit1(prices []int) int {

	var count int                        // 初始化利润
	for i := 0; i < len(prices)-1; i++ { // 开始遍历，注意下标问题
		if prices[i+1] > prices[i] { // 判断，当前下标后一个元素比较当前元素大
			count += prices[i+1] - prices[i] // 下一个元素减当前元素，利润累加
		}
	}
	return count
}
