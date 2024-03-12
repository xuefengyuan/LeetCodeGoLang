package _020_Valid_Parentheses

import "testing"

/**
 * @Author: darry
 * @Desc:
 * @Date: 2024/3/12 21:51
 */

func Test_isValid(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {"()", args{s: "()"}, true},
        {"()[]{}", args{s: "()[]{}"}, true},
        {"{([])}", args{s: "{([])}"}, true},
        {"(]", args{s: "(]"}, false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isValid(tt.args.s); got != tt.want {
                t.Errorf("isValid() = %v, want %v", got, tt.want)
            }
        })
    }
}
