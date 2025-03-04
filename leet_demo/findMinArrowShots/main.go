package main

import (
	"fmt"
	"sort"
)

// 452 用最少数量的箭引爆气球
// 有一些球形气球贴在一堵用 XY 平面表示的墙面上。墙面上的气球记录在整数数组 points ，其中points[i] = [xstart, xend] 表示水平直径在 xstart 和 xend之间的气球。你不知道气球的确切 y 坐标。
//
// 一支弓箭可以沿着 x 轴从不同点 完全垂直 地射出。在坐标 x 处射出一支箭，若有一个气球的直径的开始和结束坐标为 xstart，xend， 且满足  xstart ≤ x ≤ xend，则该气球会被 引爆 。可以射出的弓箭的数量 没有限制 。 弓箭一旦被射出之后，可以无限地前进。
//
// 给你一个数组 points ，返回引爆所有气球所必须射出的 最小 弓箭数 。
func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))

	points = [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(findMinArrowShots(points))

	points = [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(findMinArrowShots(points))

}

// 排序+贪心算法
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 按照每个气球的右边界升序排序
	sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })

	// 初始化射箭次数
	result := 1
	// 初始化第一次射出箭的位置
	arrowPos := points[0][1]

	for _, point := range points {
		// 如果当前气球的起始位置(左边界)大于箭的位置，需要发射新箭
		if point[0] > arrowPos {
			// 射箭次数++
			result++
			// 更新射箭位置为当前气球的右边界
			arrowPos = point[1]
		}
	}

	return result
}
