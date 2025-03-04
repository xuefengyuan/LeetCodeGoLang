package main

import "fmt"

// 35.搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
// 请必须使用时间复杂度为 O(log n) 的算法。
func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 5))

	nums = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 2))

	nums = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 7))

	nums = []int{1, 3}
	fmt.Println(searchInsert(nums, 2))

}

// 二分查找
func searchInsert(nums []int, target int) int {
	// 初始化左右指针,-1
	left := 0
	right := len(nums) - 1

	// 二分查找,条件是左指针小于等于右指针
	for left <= right {
		// 计算中间位置
		mid := left + (right-left)/2
		// 判断是不是找到了目标值
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 目标值在右半部分
			left = mid + 1
		} else {
			// 目标值在左半部分
			right = mid - 1
		}
	}

	// 返回插入位置
	return left
}
