package main

import "fmt"

func main()  {
    isRes := isPalindrome(232)
    fmt.Println("is Res 1= ",isRes)

    isRes = isPalindrome(-121)
    fmt.Println("is Res 2= ",isRes)

    isRes = isPalindrome(10)
    fmt.Println("is Res 3= ",isRes)

    isRes = isPalindrome(898)
    fmt.Println("is Res 4= ",isRes)

    isRes = isPalindrome(1475741)
    fmt.Println("is Res 5= ",isRes)
}


func isPalindrome(x int) bool {
    y := x
    var num int
    for y != 0 {
        num = num * 10 + y % 10
        y = y / 10
    }
    if x < 0 || num != x {
        return false
    }
    return true
}
