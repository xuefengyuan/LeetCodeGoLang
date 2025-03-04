package main

import "fmt"

// 205 同构字符串
// 给定两个字符串 s 和 t ，判断它们是否是同构的。
//
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
//
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身
func main() {
	s, t := "egg", "add"
	ret := isIsomorphic(s, t)
	fmt.Println(ret)

	s, t = "foo", "bar"
	ret = isIsomorphic(s, t)
	fmt.Println(ret)

	s, t = "paper", "title"
	ret = isIsomorphic(s, t)
	fmt.Println(ret)

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 创建两个映射表，用于存储字符之间的映射关系
	SM := map[byte]byte{}
	tM := map[byte]byte{}

	for i := range s {

		cs, ct := s[i], t[i]

		// 判断两边的映射关系是否能对得上
		if SM[cs] > 0 && SM[cs] != ct || tM[ct] > 0 && tM[ct] != cs {
			return false
		}

		SM[cs] = ct
		tM[ct] = cs
	}

	return true
}
