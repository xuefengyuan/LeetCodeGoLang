package main

import "fmt"

// 罗马数字转整数
// 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
//
//	字符          数值
//	I             1
//	V             5
//	X             10
//	L             50
//	C             100
//	D             500
//	M             1000
func main() {
	s := "III"
	toInt := romanToInt(s)
	fmt.Println(toInt)

	s = "IV"
	toInt = romanToInt(s)
	fmt.Println(toInt)

	s = "IX"
	toInt = romanToInt(s)
	fmt.Println(toInt)

	s = "LVIII"
	toInt = romanToInt(s)
	fmt.Println(toInt)

	s = "MCMXCIV"
	toInt = romanToInt(s)
	fmt.Println(toInt)

}

func romanToInt(s string) int {

	// 定义罗马数字和整数的映射关系
	ri := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	// 定义结果，获取字符串长度，为0就返回
	var res int
	n := len(s)
	if n == 0 {
		return res
	}

	for i := range s {
		v := ri[s[i]]                  // 获取下标获取罗马字符串元素对应的整数
		if i < n-1 && v < ri[s[i+1]] { // 判断当前下标小于整个长度减1，并且下标元素对应的整数小于后一位
			res -= v // 结果减去当前下标元素的整数
			continue
		} /*else {
			res += v // 结果加上当前下标元素整数
		}*/
		res += v // 结果加上当前下标元素整数
	}
	return res
}
