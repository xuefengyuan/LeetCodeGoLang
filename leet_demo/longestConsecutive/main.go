package main

import "fmt"

// 128 最长连续序列
func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))

	nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
}

func longestConsecutive(nums []int) int {

	// 小于边界长度，直接返回
	if len(nums) < 1 {
		return 0
	}
	//使用哈希集合 numSet 来存储数组中的所有数字
	numSet := make(map[int]bool)

	for _, num := range nums {
		// 设置哈希，存在的数字设置为true
		numSet[num] = true
	}
	result := 0
	// 开始遍历
	for num := range numSet {
		// 判断当前数字 num 是否是一个 连续序列的起点。如果 num-1 不在哈希集合中，说明当前数字是一个新序列的开始
		if !numSet[num-1] {
			currentNum := num
			currentLenght := 1
			for numSet[currentNum+1] {
				currentNum++
				currentLenght++
			}

			if result < currentLenght {
				result = currentLenght
			}
		}
	}

	return result
}
