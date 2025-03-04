package main

import "fmt"

// 33 搜索排序数组
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
// 你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。
func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))

	fmt.Println(search(nums, 5))

	fmt.Println(search(nums, 3))

	nums = []int{1}
	fmt.Println(search(nums, 0))

}

// 二分查找
func search(nums []int, target int) int {

	// 初始化左右指针，右指针是数组的长度-1
	left := 0
	right := len(nums) - 1

	// 进行二分查找，目标在左边就要缩减右指针，目标在右边就要增加左指针
	for left <= right {
		// 计算中间值
		mid := left + (right-left)/2
		// 如果找到目标，返回目标所有的下标
		if nums[mid] == target {
			return mid
		}
		// 判断下左半部分是不是有序
		if nums[left] <= nums[mid] {
			// 检查目标是不是在左半部分（左边下标值和中间下标值是不是在查找的区间内）
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				// 目标在右半部分
				left = mid + 1
			}
		} else { // 右半部分是有序
			// 检查目标是不是在右半部分（右半部分的中间下标值和右下标值是不是在区间范围内）
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				// 目标在左半部分
				right = mid - 1
			}
		}
	}
	return -1
}
