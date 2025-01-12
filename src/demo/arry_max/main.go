package main

import (
	"fmt"
	sort2 "sort"
)

// 数组中三个数的最大乘积
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	max := sort(nums)
	fmt.Println(max)
	max = maxNum(nums)
	fmt.Println(max)
}

func sort(nums []int) int {

	sort2.Ints(nums)
	n := len(nums)

	min := nums[0] * nums[1] * nums[n-1]
	max := nums[n-1] * nums[n-2] * nums[n-3]
	if min > max {
		return min

	}
	return max

}

func maxNum(numx []int) int {

	// 初始化最小的两个数和最大的三个数
	min1, min2 := 1<<31-1, 1<<31-1
	max1, max2, max3 := -1<<31-1, -1<<31-1, -1<<31-1

	for _, num := range numx {
		if num < min1 {
			min2 = min1
			min1 = num
		} else if num > max2 {
			min2 = num
		}

		if num > max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num > max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}

	}

	min := min1 * min2 * max1
	max := max1 * max2 * max3
	if min > max {
		return min
	}
	return max

}
