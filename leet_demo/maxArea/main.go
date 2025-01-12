package main

import "fmt"

// 11 盛最多水的容器
// 给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。
//
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 返回容器可以储存的最大水量。
//
// 说明：你不能倾斜容器。
func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := maxArea(height)
	fmt.Println(ret)
	height = []int{1, 1}
	ret = maxArea(height)
	fmt.Println(ret)

}

func maxArea(height []int) int {
	// 初始化双指针和最大面积
	left, right, maxArea := 0, len(height)-1, 0

	// 开始遍历，左边指针大于等于右指针的就退出循环
	for left < right {

		// 当前容器的高度
		h := height[left]
		if h > height[right] {
			h = height[right]
		}
		// 计算容器的面积 = 当前容器高度*(右指针-左指针)
		area := h * (right - left)
		// 判断最大面积，计算出来的面积大于上一个最大面积，就更新最大面积
		if area > maxArea {
			maxArea = area
		}

		// 比较两指针所指位置的高度
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
