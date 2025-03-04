package main

import (
	"fmt"
	"math"
)

// 209 长度最小的子数组
// 给定一个含有 n 个正整数的数组和一个正整数 target 。
//
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
func main() {
	target := 7
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(target, nums))

	target = 4
	nums = []int{1, 4, 4}
	fmt.Println(minSubArrayLen(target, nums))

	target = 11
	nums = []int{1, 1, 1, 1, 1, 1, 1, 1}
	fmt.Println(minSubArrayLen(target, nums))
}

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 初始化最小长度为最大整数值
	minLenght := math.MaxInt32
	// 当前窗口的元素和
	sum := 0
	// 滑动窗口的起始位置
	start := 0

	// 遍历数组的每个元素
	for end := 0; end < n; end++ {
		// 增加当前元素到窗口中
		sum += nums[end]

		// 当窗口内的和大于等于 target 时，开始收缩窗口
		for sum >= target {
			// 更新最小长度,最小长度大于end-start+1就更新最小长度
			if minLenght > end-start+1 {
				minLenght = end - start + 1
			}
			// 缩小窗口：移除窗口的左边元素
			sum -= nums[start]
			start++
		}
	}

	// 如果没有找到符合条件的子数组，返回 0
	if minLenght == math.MaxInt32 {
		return 0
	}

	return minLenght
}
