package _014_Longest_Common_Prefix

import "testing"

/**
 * @Author: darry
 * @Desc:
 * @Date: 2024/3/11 22:27
 */

func Test_longestCommonPrefix(t *testing.T) {
    type args struct {
        strs []string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {"fl", args{strs: []string{"flower", "flow", "flight"}}, "fl"},
        {"空字符串", args{strs: []string{"dog", "racecar", "car"}}, ""},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := longestCommonPrefix(tt.args.strs); got != tt.want {
                t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
            }
        })
    }
}
