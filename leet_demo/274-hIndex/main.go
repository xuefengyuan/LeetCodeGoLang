package main

import (
	"fmt"
	"sync"
	"time"
)

// 274 H指数
// 给你一个整数数组 citations ，其中 citations[i] 表示研究者的第 i 篇论文被引用的次数。计算并返回该研究者的 h 指数。
//
// 根据维基百科上 h 指数的定义：h 代表“高引用次数” ，一名科研人员的 h 指数 是指他（她）至少发表了 h 篇论文，并且 至少 有 h 篇论文被引用次数大于等于 h 。如果 h 有多种可能的值，h 指数 是其中最大的那个。
func main1() {
	citations := []int{3, 0, 6, 1, 5}
	ret := hIndex(citations)
	fmt.Println(ret)
	citations = []int{1, 3, 1}
	ret = hIndex(citations)
	fmt.Println(ret)
}
func main2() {
	a := []int{1, 2, 3}
	b := a[:2]

	b = append(b, 4)
	fmt.Println("a1", a) // 1,2,4,3

	b = append(b, 5)
	fmt.Println(a) // 1,2,5,4,3

	b[0] = 10
	fmt.Println(a) // 1,2,10,4,3
}

var count int

func increment() { count++ }

func main3() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	time.Sleep(time.Second)
	fmt.Println(count)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

// 计数排序
func hIndex(citations []int) int {
	// 论文的总数
	n := len(citations)
	// 创建一个长度为总数长度+1的桶数组，每个值表示一篇论文的引用次数
	count := make([]int, n+1)
	// 统计引用次数，c为单个引用次数
	for _, c := range citations {
		// 如果引用次数大于或等于N，引用次数≥n的归到count[n]
		if c >= n {
			count[n]++
		} else {
			// 否则，归入对应的桶，计入对应索引
			count[c]++
		}
	}

	// 累计论文数
	sum := 0
	// 倒序往前遍历
	for i := n; i >= 0; i-- {
		// 累加当前和上次引用次数的论文
		sum += count[i]
		// 如果累加的论文数大于等于当前引用次数
		if sum >= i {
			// 返回指数，就是引用次数
			return i
		}
	}
	// 默认返回0，没有引用
	return 0
}
