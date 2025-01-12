package main

import "fmt"

// 169 多元数组
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
func main() {

	nums := []int{3, 2, 3}
	result := majorityElement(nums)
	fmt.Println(result)

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	result = majorityElement(nums)
	fmt.Println(result)

}

// 投票算法
func majorityElement(nums []int) int {

	count := 0
	var data int
	for _, num := range nums {
		if count == 0 {
			data = num
		}

		if num == data {
			count++
		} else {
			count--
		}

	}
	return data
}
