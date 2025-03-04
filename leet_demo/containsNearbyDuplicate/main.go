package main

import "fmt"

// 219 存在重复元素
// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false
func main() {
	nums, k := []int{1, 2, 3, 1}, 3
	fmt.Println(containsNearbyDuplicate(nums, k))

	nums, k = []int{1, 0, 1, 1}, 2
	fmt.Println(containsNearbyDuplicate(nums, k))

	nums, k = []int{1, 2, 3, 1, 2, 3}, 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}

// 滑动窗口+哈希表
func containsNearbyDuplicate(nums []int, k int) bool {

	// 创建一个哈希表，用于存储元素及其最近的下标
	hmap := make(map[int]int)
	// 遍历数组
	for i, num := range nums {

		// 如果当前元素已经存在于哈希表中，检查条件 abs(i - j) <= k
		if idx, ok := hmap[num]; ok && i-idx <= k {
			return true
		}

		// 更新哈希表中当前元素的下标
		hmap[num] = i
	}

	// 如果没有找到满足条件的元素对，返回 false
	return false
}
