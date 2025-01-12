package main

import "fmt"

// 55、 跳跃游戏
// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。
//
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
func main() {
	nums := []int{2, 3, 1, 1, 4}
	ret := canJump(nums)
	fmt.Println(ret)

	nums = []int{3, 2, 1, 0, 4}
	ret = canJump(nums)
	fmt.Println(ret)
}

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReach {
			return false
		}
		maxReach = maxReach
		if maxReach < i+nums[i] {
			maxReach = i + nums[i]
		}
	}

	return true
}
