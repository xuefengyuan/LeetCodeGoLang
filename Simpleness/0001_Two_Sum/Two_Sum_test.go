package _001_Two_Sum

/**
 * @Author: darry
 * @Desc:
 * @Date: 2024/3/7 20:05
 */

import (
    "reflect"
    "testing"
)

func TestTwoSum(t *testing.T) {
    nums := []int{2, 7, 11, 15}
    target := 9
    expected := []int{0, 1}

    result := TwoSum(nums, target)
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("Expected %v, but got %v", expected, result)
    }
}
