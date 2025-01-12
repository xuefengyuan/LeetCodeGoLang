package main

import "fmt"

// x的平方根，二分查找和牛顿迭代方式

func main() {
	search := binarySearch(6)
	fmt.Println(search)
	//search = A(25)
	search = newTon(6)
	fmt.Println(search)
}

// 二分查找
func binarySearch(x int) int {
	index := -1
	l := 0
	r := x

	for l <= r {
		mid := l + (r-l)/2
		if mid*mid <= x {
			index = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return index
}

// 牛顿迭代
func newTon(x int) int {

	return int(sqrt(x, float64(x)))

}

func sqrt(x int, i float64) float64 {
	res := (i + float64(x)/i) / 2
	if res == i {
		return i

	} else {
		return sqrt(x, res)
	}
}
