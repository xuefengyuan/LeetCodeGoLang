package main

import (
	"fmt"
	"strings"
)

// 6 Z字形变换
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	ret := convert(s, numRows)
	fmt.Println(ret)

	s = "PAYPALISHIRING"
	numRows = 4
	ret = convert(s, numRows)
	fmt.Println(ret)

	s = "A"
	numRows = 1
	ret = convert(s, numRows)
	fmt.Println(ret)

}

func convert(s string, numRows int) string {
	// 特殊情况，直接返回
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	// 创建一个rows，存储每一行字符串
	rows := make([]strings.Builder, numRows)
	curRow := 0        // 当前处理行
	goingDown := false // 方向标志，当前是向上还是向下
	for _, char := range s {
		rows[curRow].WriteRune(char)
		if curRow == 0 || curRow == numRows-1 { // 如果是顶部或者询问，反转方向
			goingDown = !goingDown
		}
		if goingDown { // 根据方向标志，调整对应所行位置
			curRow++
		} else {
			curRow--
		}
	}

	// 合并成一个新的字符串
	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}
	return result.String()
}
