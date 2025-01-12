package main

import "fmt"

// 42 接雨水
//
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水
func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	ret := trap(height)
	fmt.Println(ret)

	height = []int{4, 2, 0, 3, 2, 5}
	ret = trap(height)
	fmt.Println(ret)
}

// 双指针
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	// 定义双指针，分别是从左和从右开始
	left, right := 0, n-1
	// 分别记录从左侧和右侧到当前位置的最大高度，最终接水的数量
	leftMax, rightMax, woter := 0, 0, 0
	// 左边没跟右边相遇的时候就继续循环
	for left < right {
		// 比较两侧的高低，处理较低的一侧，左边小于右边就进入左边的计算
		if height[left] < height[right] {
			// 数组左边下标高度大于或等于左侧最大值
			if height[left] >= leftMax {
				// 更新左边最大值 = 数组当前左边下标的高度
				leftMax = height[left]
			} else {
				// 当前高度可以接水
				// 接水量=左边最大值-数组左边下标的高度
				woter += leftMax - height[left]
			}
			// 指针向右移动
			left++
		} else { // 这里是左边大于右边，进入右边的计算
			// 数组右边下标高度大于等于右边最大高度
			if height[right] >= rightMax {
				// 更新右边最大值 = 数组右边下标的高度
				rightMax = height[right]
			} else {
				//  接水量 = 右边最大值-数组右边下标高度
				woter += rightMax - height[right]
			}
			// 指针向左移动
			right--
		}
	}
	return woter
}
