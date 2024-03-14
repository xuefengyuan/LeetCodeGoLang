/**
 * @Author: darry
 * @Desc: 删除有序数组中的重复项
示例 1：

输入：nums = [1,1,2]
输出：2, nums = [1,2,_]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。
示例 2：

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。
 * @Date: 2024/3/14 21:37
*/
package _026_Remove_Duplicates_from_Sorted_Array

import "fmt"

func removeDuplicates(nums []int) int {

    n := len(nums)

    if n <= 1 {
        return n
    }

    inx := 1

    for i := 1; i < n; i++ {
        if nums[i] != nums[i-1] {
            nums[inx] = nums[i]
            inx++
        }
    }
    fmt.Printf("inx : %d, nums = %+v\n", inx, nums)
    return inx
}
