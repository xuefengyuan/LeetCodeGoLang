package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 380 O(1)时间插入，删除和获取随机元素
// 实现RandomizedSet 类：
//
// RandomizedSet() 初始化 RandomizedSet 对象
// bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
// bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
// int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
// 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1)
func main() {

	rand.NewSource(time.Now().UnixNano())
	//rand.Seed(time.Now().UnixNano()) // 随机数种子
	// 创建集合
	rc := Constructor()

	// 插入元素
	rc.Insert(1)
	rc.Insert(2)
	rc.Insert(3)

	// 获取随机元素
	fmt.Println(rc.nums)
	fmt.Println("获取删除前的随机元素：", rc.GetRandom())

	// 删除元素
	rc.Remove(2)

	// 获取删除后的随机元素
	fmt.Println(rc.nums)
	fmt.Println("获取删除后的随机元素：", rc.GetRandom())
}

type RandomizedSet struct {
	nums   []int
	indexs map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indexs[val]; ok {
		return false
	}
	rs.indexs[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)

	return true
}

func (rs *RandomizedSet) Remove(val int) bool {

	id, ok := rs.indexs[val]
	if !ok {
		return false
	}

	last := len(rs.nums) - 1
	rs.nums[id] = rs.nums[last]
	rs.indexs[rs.nums[id]] = id
	rs.nums = rs.nums[:last]
	delete(rs.indexs, val)

	return true
}

func (rs *RandomizedSet) GetRandom() int {

	return rs.nums[rand.Intn(len(rs.nums))]
}
