package main

import (
	"fmt"
	"strconv"
)

// 228 汇总区间
// 给定一个  无重复元素 的 有序 整数数组 nums 。
//
// 返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。
//
// 列表中的每个区间范围 [a,b] 应该按如下格式输出：
// "a->b" ，如果 a != b
// "a" ，如果 a == b
func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(nums))

	nums = []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	start, end := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == end+1 {
			// 如果当前数字和前一个数字连续，更新区间终点
			end = nums[i]
		} else {
			// 当前数字和前一个数字不连续，将当前区间加入结果
			s := strconv.Itoa(start)
			if start != end {
				s = strconv.Itoa(start) + "->" + strconv.Itoa(end)
			}

			result = append(result, s)
			// 开始一个新区间
			start, end = nums[i], nums[i]
		}
	}
	// 将最后一个区间加入结果
	s := strconv.Itoa(start)
	if start != end {
		s = strconv.Itoa(start) + "->" + strconv.Itoa(end)
	}
	result = append(result, s)

	return result
}

func summaryRanges1(nums []int) []string {
	result := make([]string, 0)

	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {
		}

		s := strconv.Itoa(nums[left])

		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}

		result = append(result, s)

	}

	return result
}
