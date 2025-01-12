package main

import "fmt"

// 移除元素
//
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	element := removeElement1(nums, val)
	fmt.Println(element)
	fmt.Println(nums)

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	element = removeElement1(nums, val)
	fmt.Println(element)
	fmt.Println(nums)
}

// 双指针
func removeElement(nums []int, val int) int {
	left := 0
	for _, num := range nums {
		if num != val {
			nums[left] = num
			left++
		}
	}

	return left
}

// 双指针优化
func removeElement1(nums []int, val int) int {
	// 初始化两个指针，一个是从nums的0开始，一个是从nums的长度开始
	left, right := 0, len(nums)
	for left < right { // 左指针小于右指针循环继续
		if nums[left] == val { // 判断左指针下标的元素等于传入的值
			nums[left] = nums[right-1] // 需要替换元素，把nums右指针-1的值，赋值给左指针下标
			right--                    // 右指针往前移
		} else {
			left++ // 左指针下标的元素不等于传入的元素，保留元素，左指针往后移
		}
	}
	// 返回不等于val的元素个数
	return left
}
