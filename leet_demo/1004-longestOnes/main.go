package main

import "fmt"

// 1004.最大连续1的个数 III
// 给定一个二进制数组 nums 和一个整数 k，假设最多可以翻转 k 个 0 ，则返回执行操作后 数组中连续 1 的最大个数 。
func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(nums, 2))

	//nums = []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	//fmt.Println(longestOnes(nums, 3))
}

// 滑动窗口实现
func longestOnes(nums []int, k int) int {
	// 窗口左边界
	left := 0
	// 窗口内0的数量
	zeros := 0
	// 最长1的长度
	maxLen := 0
	// 滑动窗口，按数量长度为条件判断
	for right := 0; right < len(nums); right++ {
		// 滑动的时候遇到0，窗口内的0累加
		if nums[right] == 0 {
			zeros++
		}

		// 如果0的数量超过k，收缩窗口，这里用循环一直判断
		for zeros > k {
			// 表示移出窗口的值，等于0的话，0的数量就减少
			if nums[left] == 0 {
				zeros--
			}
			// 移动左边界
			left++
		}
		// 更新最大长度
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
