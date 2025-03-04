package main

import (
	"fmt"
	"strings"
)

// 71 简化路径
// 给你一个字符串 path ，表示指向某一文件或目录的 Unix 风格 绝对路径 （以 '/' 开头），请你将其转化为 更加简洁的规范路径。
//
// 在 Unix 风格的文件系统中规则如下：
//
// 一个点 '.' 表示当前目录本身。
// 此外，两个点 '..' 表示将目录切换到上一级（指向父目录）。
// 任意多个连续的斜杠（即，'//' 或 '///'）都被视为单个斜杠 '/'。
// 任何其他格式的点（例如，'...' 或 '....'）均被视为有效的文件/目录名称。
// 返回的 简化路径 必须遵循下述格式：
//
// 始终以斜杠 '/' 开头。
// 两个目录名之间必须只有一个斜杠 '/' 。
// 最后一个目录名（如果存在）不能 以 '/' 结尾。
// 此外，路径仅包含从根目录到目标文件或目录的路径上的目录（即，不含 '.' 或 '..'）。
// 返回简化后得到的 规范路径 。
func main() {
	path := "/home/"
	fmt.Println(simplifyPath(path))

	path = "/home//foo/"
	fmt.Println(simplifyPath(path))

	path = "/home/user/Documents/../Pictures"
	fmt.Println(simplifyPath(path))
}

func simplifyPath(path string) string {
	// 分割路径
	parts := strings.Split(path, "/")
	stack := []string{}

	for _, part := range parts {
		// 跳过空字符串和 "."
		if part == "" || part == "." {
			continue
		} else if part == ".." {
			// 如果是 ".."，弹出栈顶元素（如果栈不为空）
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		// 入栈
		stack = append(stack, part)
	}

	return "/" + strings.Join(stack, "/")
}
