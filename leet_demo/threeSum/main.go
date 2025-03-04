package main

import (
	"fmt"
	"sort"
	"unsafe"
)

type Example struct {
	A int8  // 1字节
	B int64 // 8字节
	C int16 // 2字节
}

// 15 三数之和
// 给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//
// 注意：答案中不可以包含重复的三元组
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := threeSum(nums)
	fmt.Println(ret)

	nums = []int{0, 1, 1}
	ret = threeSum(nums)
	fmt.Println(ret)

	nums = []int{0, 0, 0}
	ret = threeSum(nums)
	fmt.Println(ret)

	var ex Example
	fmt.Println("Size of struct:", unsafe.Sizeof(ex)) // 输出: 16
}

func threeSum(nums []int) [][]int {
	// 用于存储结果
	result := make([][]int, 0)
	n := len(nums)

	// 对数组排序
	sort.Ints(nums)

	for i := 0; i < n-2; i++ {
		// 跳过重复的第一个数字
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 初始化双指针
		target := -nums[i]
		left, right := i+1, n-1
		// 双指针寻找两数之和
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				// 找到一个三元组
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的第二个数字
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				// 跳过重复的第三个数字
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				// 移动指针
				left++
				right--
			} else if sum < target {
				// 和小于目标值，移动左指针以增大和
				left++
			} else {
				// 和大于目标值，移动右指针以减小和
				right--
			}
		}
	}

	return result
}
