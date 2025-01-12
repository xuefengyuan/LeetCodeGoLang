package main

import "fmt"

// 80 删除有序数组中的重复项 II
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	ret := removeDuplicates(nums)
	fmt.Println(ret, nums)

	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	ret = removeDuplicates(nums)
	fmt.Println(ret, nums)
}

func removeDuplicates(nums []int) int {
	// 因为最少都要保留两个元素，所以小于等于2的时候直接返回
	n := len(nums)
	if n <= 2 {
		return n
	}

	k := 2                   // 初始化下标从2开始
	for i := 2; i < n; i++ { // 遍历的下标也是从2开始
		if nums[i] != nums[k-2] { // 判断要是相等的话，元素是已经出现了，跳过，不相等的话，保留当前元素，放在k下标的位置，k下标加1
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
