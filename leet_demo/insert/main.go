package main

import "fmt"

// 57 插入区间
// 给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，其中 intervals[i] = [starti, endi] 表示第 i 个区间的开始和结束，并且 intervals 按照 starti 升序排列。同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始和结束。
//
// 在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以合并区间）。
//
// 返回插入之后的 intervals。
//
// 注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。
func main() {
	intervals, newInterval := [][]int{{1, 3}, {6, 9}}, []int{2, 5}
	fmt.Println(insert(intervals, newInterval))

	intervals, newInterval = [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}
	fmt.Println(insert(intervals, newInterval))

}

func insert(intervals [][]int, newInterval []int) [][]int {

	result := [][]int{}
	if len(intervals) == 0 {
		result = append(result, newInterval)
		return result
	}
	inserted := false

	for _, interval := range intervals {
		// 当前区间在 newInterval 左侧，且不重叠
		if interval[1] < newInterval[0] {
			result = append(result, interval)
		} else if interval[0] > newInterval[1] {
			// 当前区间在 newInterval 右侧，且不重叠
			if !inserted {
				result = append(result, newInterval)
				inserted = true
			}
			result = append(result, interval)
		} else {
			// 当前区间与 newInterval 重叠，合并区间
			newInterval[0] = min(newInterval[0], interval[0])
			newInterval[1] = max(newInterval[1], interval[1])
		}
	}
	// 如果 newInterval 未插入，则加入结果
	if !inserted {
		result = append(result, newInterval)
	}

	return result
}

// 返回两个数的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 返回两个数的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
