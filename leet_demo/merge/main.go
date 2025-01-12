package main

import (
	"fmt"
	"sort"
)

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2，另有两个整数 m 和 n ，分别表示 nums1 和 nums2 中的元素数目。
//
// 请你 合并 nums2 到 nums1 中，使合并后的数组同样按 非递减顺序 排列。

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge2(nums1, m, nums2, n)
	fmt.Printf("%v\n", nums1)
}

// 合并后排序
func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

// 双指针
func merge1(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}

	}
	copy(nums1, sorted)
}

// 逆向双指针
func merge2(nums1 []int, m int, nums2 []int, n int) {

	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		// 下面的--都是指针往前移动
		if p1 == -1 { // nums1处理完了，取nums2里面的全部元素
			cur = nums2[p2]
			p2--
		} else if p2 == -1 { // nums2处理完了，取nums1里面的全部元素
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] { // nums1的元素比nums2大，要放到后面
			cur = nums1[p1]
			p1--
		} else { // nums2的元素大，要放到后面
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur // 替换下标中的元素
	}
}
