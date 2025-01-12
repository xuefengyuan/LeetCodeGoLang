package main

import "fmt"

func main() {
	// 测试用例

	//num1 := "987654321987654321987654321"
	//num2 := "123456789123456789123456789"

	num1 := "4321907154021"
	num2 := "89123456059"
	result := addlargeNum(num1, num2)
	fmt.Println("Result:", result) // 输出结果
	//var num = int8(110)
	//fmt.Println(num)
	//var i, j, m, n = 99, 3.1415926, 'M', "this is a string"
}

func addlargeNum(num1, num2 string) string {

	result := "" // 结果
	carry := 0   //是否有进位

	// 从个位开始相加
	i, j := len(num1)-1, len(num2)-1

	for i >= 0 || j >= 0 || carry > 0 {

		dit1 := 0
		dit2 := 0
		if i >= 0 {
			dit1 = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			dit2 = int(num2[j] - '0')
			j--
		}
		sum := dit1 + dit2 + carry
		carry = sum / 10
		fmt.Println("carry = ", carry)
		result = string(sum%10+'0') + result

	}
	return result
}
