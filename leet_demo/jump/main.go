package main

import "fmt"

// 45、跳跃游戏II
// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。
//
// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处
func main() {
	nums := []int{2, 3, 1, 1, 4}
	ret := jump(nums)
	fmt.Println(ret)

	nums = []int{3, 2, 1, 0, 4}
	ret = jump(nums)
	fmt.Println(ret)
}

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	jumps, farthest, currentEnd := 0, 0, 0
	for i := 0; i < n-1; i++ {
		farthest = farthest
		if farthest < i+nums[i] {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}

	return jumps

}
