package main

import (
	"fmt"
)

// 134 加油站
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
//
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
//
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的
func main() {
	gas, cost := []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}
	ret := canCompleteCircuit(gas, cost)
	fmt.Println(ret)

	gas, cost = []int{2, 3, 4}, []int{3, 4, 3}
	ret = canCompleteCircuit(gas, cost)
	fmt.Println(ret)
}

// 双指针+贪心算法
// gas 每个加油站可以提供的汽油量
// cost 从当前加油站到下一个加油站所需要的油量
func canCompleteCircuit(gas []int, cost []int) int {
	// 初始化总油量余量，当前油量，起点索引
	totalTank, currTank, start := 0, 0, 0
	// 遍历所有加油站
	for i := 0; i < len(gas); i++ {
		// 更新总油量余量
		totalTank += gas[i] - cost[i]
		//更新当前油量余量
		currTank += gas[i] - cost[i]
		// 当前加油站走不到下一个加油站
		if currTank < 0 {
			// 从下一个加油站开始
			start = i + 1
			// 重置当前油量
			currTank = 0
		}
	}
	// 检查总油量是不是负数，负数返回-1
	if totalTank < 0 {
		return -1
	}

	return start
}

// gas 每个加油站可以提供的汽油量
// cost 从当前加油站到下一个加油站所需要的油量
func canCompleteCircuit1(gas []int, cost []int) int {
	total, curr := 0, 0
	start := 0

	for i := 0; i < len(gas); i++ {
		total += gas[i] - cost[i]
		curr += gas[i] - cost[i]
		if curr < 0 {
			start = i + 1
			curr = 0
		}
	}
	if total >= 0 {
		return start
	}

	return -1
}
