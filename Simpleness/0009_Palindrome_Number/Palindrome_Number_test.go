package _009_Palindrome_Number

import "testing"

/**
 * @Author: darry
 * @Desc:
 * @Date: 2024/3/8 22:17
 */

func Test_isPalindrome(t *testing.T) {
    type args struct {
        x int
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {"12721", args{x: 12721}, true},
        {"-121", args{x: -121}, false},
        {"450", args{x: 450}, false},
        {"0", args{x: 0}, true},
        {"374686473", args{x: 374686473}, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isPalindrome(tt.args.x); got != tt.want {
                t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
            }
        })
    }
}

func Test_isPalindrome1(t *testing.T) {
    type args struct {
        x int
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {"12721", args{x: 12721}, true},
        {"-121", args{x: -121}, false},
        {"450", args{x: 450}, false},
        {"0", args{x: 0}, true},
        {"374686473", args{x: 374686473}, true},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isPalindrome1(tt.args.x); got != tt.want {
                t.Errorf("isPalindrome1() = %v, want %v", got, tt.want)
            }
        })
    }
}
