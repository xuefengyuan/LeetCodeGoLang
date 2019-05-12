package main

import "fmt"

func main()  {
    nums := []int{2,7,11,15}
    sum := twoSum(nums, 26)

    fmt.Println("sun = ",sum)

}

func twoSum(nums []int, target int) []int {
    for i := 0; i<len(nums); i++{
        for j := i+1; j < len(nums); j++{
            if nums[j] == target - nums[i]{
                return []int{i,j}
            }
        }
    }
    return nil
}

