package main

import (
    "fmt"
    "math"
)

func main()  {
    res := reverse(123)
    fmt.Println("res1 = ",res)

    res = reverse(-123)
    fmt.Println("res2 = ",res)

    res = reverse(120)
    fmt.Println("res3 = ",res)
}

func reverse(x int) int {
    var num int
    for x != 0{ // 直到x等于0，跳出循环
        num = num*10 + x%10
        x = x/10
    }
    // 题目要求其数值范围是 [−2^31,  2^31 − 1]。如果反转后的整数溢出，则返回 0。
    if num > math.MaxInt32 || num < math.MinInt32{
        return 0
    }

    return num
}
