package _026_Remove_Duplicates_from_Sorted_Array

import "testing"

/**
 * @Author: darry
 * @Desc:
 * @Date: 2024/3/14 21:49
 */

func Test_removeDuplicates(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {"2", args{nums: []int{1, 1, 2}}, 2},
        {"5", args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := removeDuplicates(tt.args.nums); got != tt.want {
                t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
            }
        })
    }
}
