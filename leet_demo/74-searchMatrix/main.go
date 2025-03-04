package main

import "fmt"

// 74. 搜索二维矩阵
// 给你一个满足下述两条属性的 m x n 整数矩阵：
//
// 每行中的整数从左到右按非严格递增顺序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// 给你一个整数 target ，如果 target 在矩阵中，返回 true ；否则，返回 false 。
func main() {
	// 测试用例
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	target := 3
	fmt.Println(searchMatrix(matrix, target)) // 输出: true

	target = 13
	fmt.Println(searchMatrix(matrix, target)) // 输出: false
}

// 单次二分查找
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	// 获取行数和列数，m是行，n是列
	m, n := len(matrix), len(matrix[0])
	// 二分查找范围，左边从0开始，右边是计算矩阵大小-1
	left, right := 0, m*n-1
	// 单次二分查找
	for left <= right {
		// 计算中间索引
		mid := left + (right-left)/2
		// 映射到行号
		row := mid / n
		// 映射到列
		col := mid % n
		// 找到目标
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			// 目标在右半部分
			left = mid + 1
		} else {
			// 目标在左半部分
			right = mid - 1
		}
	}
	return false
}
