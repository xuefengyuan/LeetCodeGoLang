package _013_Roman_to_Integer

import "testing"

/**
 * @Author: darry
 * @Desc:
 * @Date: 2024/3/9 20:47
 */

func Test_romanToInt(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {"III", args{s: "III"}, 3},
        {"LVIII", args{s: "LVIII"}, 58},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := romanToInt(tt.args.s); got != tt.want {
                t.Errorf("romanToInt() = %v, want %v", got, tt.want)
            }
        })
    }
}
