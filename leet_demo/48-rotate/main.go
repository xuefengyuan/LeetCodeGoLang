package main

import "fmt"

// 48. 旋转图像
// 给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
//
// 你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。
func main() {
	// 测试用例
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("= = = = = = = = = = = = = =")
	matrix = [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	rotate(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("= = = = = = = = = = = = = =")
	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// 翻转法
func rotate(matrix [][]int) {
	// 获取矩阵的边长
	n := len(matrix)

	// 第一步：沿主对角线翻转 (i,j) 与 (j,i) 交换
	for i := 0; i < n; i++ {
		// 内层循环，遍历列索引 j，下标从 i 开始到 n-1，只处理对角线以上部分
		for j := i; j < n; j++ {
			fmt.Println("i = ", i, "j = ", j)
			// 交换 matrix[i][j] 和 matrix[j][i]，完成沿主对角线的翻转
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 第二步：沿垂直中线翻转 (i,j) 与 (i,n-1-j) 交换
	for i := 0; i < n; i++ {
		// 内层循环，遍历列索引 j，从 0 到 n/2-1，只处理每行左半部分
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
