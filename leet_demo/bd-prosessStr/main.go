package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(proStr("ab12345"))
	fmt.Println(proStr("a1b2c3"))
	fmt.Println(proStr("aabbcc1122334"))
	fmt.Println(proStr("aabbcc"))
	fmt.Println(proStr("1122334"))

}

func proStr(s string) string {
	for {
		find := false

		for i := 0; i < len(s)-1; i++ {
			if unicode.IsLetter(rune(s[i])) && unicode.IsDigit(rune(s[i+1])) {
				s = s[:i] + s[i+2:]
				find = true
				break
			}
		}
		if !find {
			break
		}
	}
	return s
}
