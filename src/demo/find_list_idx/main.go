package main

import "fmt"

// 寻找数组中心下标，中心下标不存在，返回-1

func main() {
	// 测试用例
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println("数组:", nums)
	fmt.Println("中心下标:", pivotIndex(nums)) // 预期结果：3

	nums = []int{1, 2, 3}
	fmt.Println("数组:", nums)
	fmt.Println("中心下标:", pivotIndex(nums)) // 预期结果：-1
}

func pivotIndex(nums []int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0

	for i, num := range nums {

		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
