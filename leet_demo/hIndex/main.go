package main

import "fmt"

// 274 H指数
// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
//
// 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。
func main() {
	citations := []int{3, 0, 6, 1, 5}
	ret := hIndex(citations)
	fmt.Println(ret)
	citations = []int{1, 3, 1}
	ret = hIndex(citations)
	fmt.Println(ret)
}

// 计数排序
func hIndex(citations []int) int {
	// 论文的总数
	n := len(citations)
	// 创建一个长度为总数长度+1的桶数组，每个值表示一篇论文的引用次数
	bucket := make([]int, n+1)

	for _, c := range citations {
		if c >= n {
			bucket[n]++ // 如果引用次数大于或等于N，归入桶
		} else {
			bucket[c]++ // 否则，归入对应的桶
		}
	}

	h := 0
	for i := n; i >= 0; i-- { // 倒序往前遍历
		h += bucket[i] // 累加当前和上次引用次数的论文
		if h >= i {    // 如果累加的论文数大于等于当前引用次数
			return i // 返回指数，就是引用次数
		}
	}
	return 0
}
