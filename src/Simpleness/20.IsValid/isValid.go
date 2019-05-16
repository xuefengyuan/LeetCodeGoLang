package main

import "fmt"

func main() {
    str := "()"
    is := isValid(str)
    fmt.Println("is 1= ",is)

    str =  "(]"
    is = isValid(str)
    fmt.Println("is 2= ",is)

    str = "([)]"
    is = isValid(str)
    fmt.Println("is 3= ",is)

    str = "{[]}"
    is = isValid(str)
    fmt.Println("is 4= ",is)

    str = "{[()]([]{})}"
    is = isValid(str)
    fmt.Println("is 5= ",is)
}

func isValid(s string) bool {

    //  特殊情况，空字符串返回true
    if len(s) == 0 {
        return true
    }
    //  配对字典
    m := map[string]string{")": "(", "]": "[", "}": "{"}
    //  栈
    var stack []string
    //  把字符串的每个字符放进栈中，每放一个就判断与前一个是不是配对的
    for i := 0; i<len(s) ; i++ {
        if len(stack) == 0 {
            stack = append(stack, string(s[i]))
        } else{
            //  判断是否配对
            //  如果是相同的话，那就去除栈的最后一个元素
            //  如果不相同的话，那就把源字符串的对应元素加进栈中
            if stack[len(stack)-1] == m[string(s[i])] {
                stack = stack[:len(stack)-1]
            } else {
                stack = append(stack, string(s[i]))
            }
        }
    }

    //  判断栈中是否没有元素
    //  是的话返回true
    //  否则返回false
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}
