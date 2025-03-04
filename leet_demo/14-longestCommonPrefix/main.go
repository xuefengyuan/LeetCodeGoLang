package main

import (
	"fmt"
)

// 14.最长公共前缀
func main() {

	strs := []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)

	strs = []string{"dog", "racecar", "car"}
	res = longestCommonPrefix(strs)
	fmt.Println(res)
}

func longestCommonPrefix(strs []string) string {
	// 处理为空的情况
	if len(strs) == 0 {
		return ""
	}
	// 取出第一个元素
	res := strs[0]
	// 然后从第2个下标开遍历剩下的元素
	for i := 1; i < len(strs); i++ {
		// 第二个循环，遍历取出的第一个元素，逐个比较每个下标的子元素
		for j := 0; j < len(res); j++ {
			// 第一个条件是判断长度数组元素的长度不够了，第二个是判断元素不相等了
			if len(strs[i]) <= j || res[j] != strs[i][j] {
				// 截取初始化的公共部分，位置下标j，就是公共部分的位置
				res = res[:j]
				// 跳出内循环，进入下一个字符串比较了
				break
			}
		}
	}

	return res
}
