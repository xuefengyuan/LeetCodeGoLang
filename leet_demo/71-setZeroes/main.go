package main

import "fmt"

// 71.矩阵置零
// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
func main() {
	// 测试用例
	//matrix := [][]int{
	//	{1, 1, 1},
	//	{1, 0, 1},
	//	{1, 1, 1},
	//}
	//setZeroes(matrix)
	//for _, row := range matrix {
	//	fmt.Println(row)
	//}
	//
	//fmt.Println("= = = = = = = = = = = = = =")
	//matrix = [][]int{
	//	{0, 1, 2, 0},
	//	{3, 4, 5, 2},
	//	{1, 3, 1, 5},
	//}
	//setZeroes(matrix)
	//for _, row := range matrix {
	//	fmt.Println(row)
	//}

	fmt.Println("= = = = = = = = = = = = = =")
	matrix := [][]int{
		{1, 0},
	}
	setZeroes(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// 标记法实现
func setZeroes(matrix [][]int) {

	//获取矩阵的尺寸,m是列，n是行
	m, n := len(matrix), len(matrix[0])
	// 第一行是不是有0的标志
	row0IsZero := false
	// 第一列是不是有0的标志
	col0IsZero := false

	// 检查第一行是不是有0
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			row0IsZero = true
			break
		}
	}

	// 检查第一列是不是有0
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			col0IsZero = true
			break
		}
	}

	// 用第一行和第一列标记其他行和列
	// 外层先遍历列，下标从1开始
	for i := 1; i < m; i++ {
		// 内层遍历行
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				// 标记行
				matrix[i][0] = 0
				// 标记列
				matrix[0][j] = 0
			}
		}
	}

	// 根据第一列标记置零行
	for i := 1; i < m; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// 根据第一行标记转零列
	for j := 1; j < n; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}

	// 处理第一行
	if row0IsZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	// 处理第一列
	if col0IsZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

}
