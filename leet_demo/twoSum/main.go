package main

import "fmt"

// 1 两数之和
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
//
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次相同的元素。
//
// 你可以按任意顺序返回答案。
func main() {
	nums, target := []int{2, 7, 11, 15}, 9
	ret := twoSum(nums, target)
	fmt.Println(ret)

	nums, target = []int{3, 2, 4}, 6
	ret = twoSum(nums, target)
	fmt.Println(ret)

	nums, target = []int{3, 3}, 6
	ret = twoSum(nums, target)
	fmt.Println(ret)

}

func twoSum(nums []int, target int) []int {
	// 定义一个map，存放元素和下标
	hmap := make(map[int]int)
	for i, num := range nums {
		// 判断，目标值减去当前元素是不是在map中，存在就返回map中的下标和当前下标 如 9-2=4
		if idx, ok := hmap[target-num]; ok {
			return []int{idx, i}
		}
		// 把当前元素放到map的key，value存放下标
		hmap[num] = i
	}
	return nil
}
