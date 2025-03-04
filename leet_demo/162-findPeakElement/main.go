package main

import "fmt"

// 162.寻找峰值
// 峰值元素是指其值严格大于左右相邻值的元素。
//
// 给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
//
// 你可以假设 nums[-1] = nums[n] = -∞ 。
//
// 你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
func main() {
	// 测试用例
	nums := []int{1, 2, 3, 1}
	fmt.Println(findPeakElement(nums)) // 输出: 2

	nums = []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums)) // 输出: 5
}

// 二分查找
func findPeakElement(nums []int) int {
	// 初始化左右指针，右指针是整个数组长度-1
	left, right := 0, len(nums)-1
	// 二分查找
	for left < right {
		// 计算中间索引
		mid := left + (right-left)/2
		// 中间索引的值，小于中间索引值+1的值，上升到右侧，说明峰值在右边
		// 中间索引下标跟+1的下标比较，找到哪峰值在哪边
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			// 左边可能有峰值
			right = mid
		}
	}
	return left
}
