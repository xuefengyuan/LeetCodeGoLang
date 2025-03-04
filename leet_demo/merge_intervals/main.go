package main

import (
	"fmt"
	"sort"
)

// 56 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间
func main() {

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))

	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))

}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 按区间的起点升序排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 存储合并后的区间
	result := [][]int{}
	// 添加第一个区间
	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		// 获取上一个区间和当前区间
		last := result[len(result)-1]
		current := intervals[i]
		// 检查是否重叠
		if last[1] >= current[0] {
			// 合并区间
			if last[1] < current[1] {
				last[1] = current[1]
			}
			// 更新最后一个区间
			result[len(result)-1] = last
		} else {
			//不重叠，直接加入结果
			result = append(result, current)
		}
	}
	return result
}
