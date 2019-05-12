package main

import "fmt"

func main() {
    number := romanToInt("III")
    fmt.Println("number 1 = ", number)

    number = romanToInt("IV")
    fmt.Println("number 2 = ", number)

    number = romanToInt("IX")
    fmt.Println("number 3 = ", number)

    number = romanToInt("LVIII")
    fmt.Println("number 4 = ", number)

    number = romanToInt("MCMXCIV")
    fmt.Println("number 5 = ", number)

    fmt.Println("==============================")

    number = romanToInt1("III")
    fmt.Println("number 6 = ", number)

    number = romanToInt1("IV")
    fmt.Println("number 7 = ", number)

    number = romanToInt1("IX")
    fmt.Println("number 8 = ", number)

    number = romanToInt1("LVIII")
    fmt.Println("number 9 = ", number)

    number = romanToInt1("MCMXCIV")
    fmt.Println("number 10 = ", number)

}

//  I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
//
//  X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
//
//  C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
//
// 示例
//  当输入是IV，那就是4，当输入是VI，那就是6
//
//  当输入是IX，那就是9，当输入是XI，那就是11
//
//  当输入是XL，那就是40，当输入是LX,那就是60
//
//  当输入是XC，那就是90，当输入是CX，那就是110
//
//  当输入是CD，那就是400，当输入是DC，那就是600
//
//  当输入是CM，那就是900，当输入是MC，那就是1100
func romanToInt(s string) int {
    res := 0
    m := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    last := 0
    for i := len(s) - 1; i >= 0; i-- {
        temp := m[s[i]] // 拿到最后一位
        sign := 1       //用于标记是减还是加
        if
        temp < last {
            //小数在大数的左边，要减去小数
            sign = -1
        }

        res += sign * temp
        last = temp
    }

    return res
}



// 按照题意，其实就是算出各个罗马数字，然后进行相加。
//
// 但是这个罗马数字有可能是1个字符代表一个数字，也有可能两个字符代表一个数字。
//
// 所以就用了两个map来记录下这两种不同的情况
//
// 1.然后我们就要对给定的字符串去取字符
//
// 2.但是当我们取到一个字符的时候，我们还得再取一个字符，用两个字符去特殊的map去查找是否存在
//
// 3.当存在的时候我们就记录下值
//
// 4.当没有的时候，我们就用第一个字符去正常的罗马数字map中找，并记录下值
func romanToInt1(s string) int {
    //特殊罗马数字
    specialRomanStringMap := map[string]int{"IV": 4, "IX": 9, "XL": 40, "XC": 90, "CD": 400, "CM": 900}
    //正常罗马数字
    romanStringMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    result := 0

    for len(s) != 0 { //对字符串循环
        if len(s) > 1 { //当长度大于1的时候，才有必要去特殊罗马数字map中查找
            chars := s[0:2] //首先我们得拿出两个字符去特殊的map中查找
            if v, ok := specialRomanStringMap[chars]; ok { //当存在的时候记录值
                result += v
                s = s[2:]
            } else { //不存在的时候去正常map中查找，并记录
                result += romanStringMap[string(s[0])]
                s = s[1:]
            }
        } else { //当字符串的长度小于等于1的就只能去正常的罗马数字map中查找
            result += romanStringMap[string(s[0])]
            s = s[1:]
        }
    }

    return result
}
