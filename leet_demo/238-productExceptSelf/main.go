package main

import "fmt"

// 238.除自身以外数组的乘积
// 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
// 题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
// 请不要使用除法，且在 O(n) 时间复杂度内完成此题。
func main() {
	nums := []int{1, 2, 3, 4}
	ret := productExceptSelf(nums)
	fmt.Println(ret)

	nums = []int{-1, 1, 0, -3, 3}
	ret = productExceptSelf(nums)
	fmt.Println(ret)
}

// 前后缀乘积法
func productExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n, n)
	// 第一个元素左侧无元素，索引 0 左侧无元素
	result[0] = 1

	// 第一步：计算左侧乘积，从下标1开始，计算每个位置左侧所有元素的乘积
	for i := 1; i < n; i++ {
		// 左侧累乘
		result[i] = result[i-1] * nums[i-1]
	}

	// 第二步：计算结合右侧乘积
	// 右侧乘积初始化
	right := 1
	// 开始遍历计算，下标从右侧开始
	for i := n - 1; i >= 0; i-- {
		// 乘以右侧乘积
		result[i] *= right
		// 更新右侧乘积
		right *= nums[i]
	}
	return result
}

func productExceptSelf1(nums []int) []int {
	// 获取数组长度
	n := len(nums)
	// 创建结果数组，初始值全为0，长度为数组长度
	answer := make([]int, n)

	// 前缀积计算值，初始化为1
	prefix := 1
	// 这里从前往后计算
	for i := 0; i < n; i++ {
		answer[i] = prefix // 更新结果，结果下标元素 = 前缀积计算值
		prefix *= nums[i]  // 更新前缀积 = 前缀积计算值*下标元素值
	}
	// 后缀积计算值，初始化为1
	suffix := 1
	// 这里从后往前计算
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix // 更新结果 = 结果下标元素*后缀积值
		suffix *= nums[i]   // 后缀积值 = 后缀积值*数组下标元素
	}

	return answer
}
