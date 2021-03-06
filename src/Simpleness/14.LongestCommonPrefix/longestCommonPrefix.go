package main

import (
    "fmt"
    "math"
)

func main() {

    strs := []string{"flower", "flow", "flight"}
    str := longestCommonPrefix(strs)
    fmt.Println(str)

    strs = []string{"dog","racecar","car"}
    str = longestCommonPrefix(strs)
    fmt.Println(str)
}

func longestCommonPrefix(strs []string) string {
    result := ""
    slen := len(strs)

    //  特殊情况，当切片中没有任何元素的时候返回""
    if slen == 0 {
        return result
    }

    //  找出最小长度的字符串、字符串长度以及索引
    minLen := math.MaxInt32
    minIndex := 0
    minLenStr := ""

    for i := 0; i < slen; i++ {
        if len(strs[i]) < minLen {
            minLen = len(strs[i])
            minIndex = 1
            minLenStr = strs[i]
        }
    }
    //  在strs去除掉长度最小的字符串
    strs = append(strs[:minIndex], strs[minIndex+1:]...)

    // 先对上面得到的最小长度的字符串进行for循环，得到每一个字符
    // 然后对源切片strs剩下的元素进行for循环，得到strs中每个元素的第一位字符，与最小长度得到的字符相比
    // 最后把相同的字符加到返回值中去

    for _, c := range minLenStr {
        for z := 0; z < len(strs); z++ {
            if string(c) == string(strs[z][0]) {
                strs[z] = strs[z][1:]
            } else {
                return result
            }
        }
        result += string(c)
    }
    return result
}
