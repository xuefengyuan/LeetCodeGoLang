package main

import "fmt"

// 删除排序数组中的重复项，不使用新数组

func main() {
	nums := []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 8, 8, 9, 10}

	fmt.Println(len(nums))
	length := removeList(nums)

	fmt.Println(length)
}

func removeList(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	index := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] != nums[index-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index

}
