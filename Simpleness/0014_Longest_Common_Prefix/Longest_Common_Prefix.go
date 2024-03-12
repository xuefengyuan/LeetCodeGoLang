/**
 * @Author: darry
 * @Desc: 最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。



示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。


提示：

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] 仅由小写英文字母组成
 * @Date: 2024/3/11 22:13
*/
package _014_Longest_Common_Prefix

func longestCommonPrefix(strs []string) string {

    if len(strs) == 0 {
        return ""
    }

    pre := strs[0]

    for i := 1; i < len(strs); i++ {
        for j := 0; j < len(pre); j++ {
            if len(strs[i]) <= j || strs[i][j] != pre[j] {
                pre = pre[:j]
                break
            }
        }
    }

    return pre
}
