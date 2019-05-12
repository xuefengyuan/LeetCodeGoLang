package main

import "fmt"

func main()  {
    res := numJewelsInStones("aA", "aAAbbbb")
    fmt.Println("res 1= ",res)

    res = numJewelsInStones("z", "ZZ")
    fmt.Println("res 2= ",res)

}

func numJewelsInStones(J string, S string) int {
    //方法：罗列出来J字符串的每一个元素去跟S字符串的每一个元素做比较，相同的情况下，次数输出+1
    var count int
    for j :=0; j<len(J);j++{
        for s :=0; s<len(S);s++{
            if J[j] == S[s]{
                count ++
            }
        }
    }

    return count

}