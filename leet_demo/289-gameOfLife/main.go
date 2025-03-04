package main

import "fmt"

// 289.生命游戏
// 根据 百度百科 ， 生命游戏 ，简称为 生命 ，是英国数学家约翰·何顿·康威在 1970 年发明的细胞自动机。
//
// 给定一个包含 m × n 个格子的面板，每一个格子都可以看成是一个细胞。每个细胞都具有一个初始状态： 1 即为 活细胞 （live），或 0 即为 死细胞 （dead）。每个细胞与其八个相邻位置（水平，垂直，对角线）的细胞都遵循以下四条生存定律：
//
// 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
// 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
// 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
// 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
// 下一个状态是通过将上述规则同时应用于当前状态下的每个细胞所形成的，其中细胞的出生和死亡是 同时 发生的。给你 m x n 网格面板 board 的当前状态，返回下一个状态。
//
// 给定当前 board 的状态，更新 board 到下一个状态。
//
// 注意 你不需要返回任何东西。
func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println("= = = = = = = = = = = = = = = = = = = = =")

	board = [][]int{{1, 1}, {1, 0}}
	gameOfLife(board)
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println("= = = = = = = = = = = = = = = = = = = = =")
}

// 状态编码法
func gameOfLife(board [][]int) {
	// 获取矩阵尺寸,m 是列,n是行
	m, n := len(board), len(board[0])

	// 第一步：遍历并标记下一代状态
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			liveCount := getLiveCount(board, i, j, m, n)
			// 判断当前是活细腻
			if board[i][j]&1 == 1 {
				// 高位设1，下一代存活
				if liveCount == 2 || liveCount == 3 {
					board[i][j] ^= 2
				}
			} else {
				// 当前是死细胞
				if liveCount == 3 {
					// 高位设1，下一代复活
					board[i][j] ^= 2
				}
			}
		}
	}

	// 第二步: 更新为下一代状态
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 右移一位,高位变为当前状态
			board[i][j] >>= 1
		}
	}
}

// 统计活邻居的数量
func getLiveCount(board [][]int, i, j, m, n int) int {
	count := 0
	// 8个方向的偏移
	diretions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	for _, dir := range diretions {
		ni, nj := i+dir[0], j+dir[1]
		if ni >= 0 && ni < m && nj >= 0 && nj < n {
			// 只取低位(当前状态)
			count += board[ni][nj] & 1
		}
	}
	return count
}
