package main

import "fmt"

// 12 整数转罗马数字
// 七个不同的符号代表罗马数字，其值如下：
//
// 符号	值
// I	1
// V	5
// X	10
// L	50
// C	100
// D	500
// M	1000
// 罗马数字是通过添加从最高到最低的小数位值的转换而形成的。将小数位值转换为罗马数字有以下规则：
// 如果该值不是以 4 或 9 开头，请选择可以从输入中减去的最大值的符号，将该符号附加到结果，减去其值，然后将其余部分转换为罗马数字。
// 如果该值以 4 或 9 开头，使用 减法形式，表示从以下符号中减去一个符号，例如 4 是 5 (V) 减 1 (I): IV ，9 是 10 (X) 减 1 (I)：IX。仅使用以下减法形式：4 (IV)，9 (IX)，40 (XL)，90 (XC)，400 (CD) 和 900 (CM)。
// 只有 10 的次方（I, X, C, M）最多可以连续附加 3 次以代表 10 的倍数。你不能多次附加 5 (V)，50 (L) 或 500 (D)。如果需要将符号附加4次，请使用 减法形式。
// 给定一个整数，将其转换为罗马数字。
func main() {
	num := 3749
	ret := intToRoman(num)
	fmt.Println(ret)

	num = 58
	ret = intToRoman(num)
	fmt.Println(ret)

	num = 1994
	ret = intToRoman(num)
	fmt.Println(ret)
}

func intToRoman(num int) string {
	// 对应1000的位数
	thousands := []string{"", "M", "MM", "MMM"}
	// 对应100的位数
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	// 对应10的位数
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	// 对应个的位数
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return thousands[num/1000] + hundreds[(num%1000)/100] + tens[(num%100)/10] + ones[num%10]
}
